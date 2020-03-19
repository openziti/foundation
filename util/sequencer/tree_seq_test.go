/*
	Copyright NetFoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package sequencer

import (
	"github.com/emirpasic/gods/trees/btree"
	"github.com/emirpasic/gods/utils"
	"github.com/pkg/errors"
	"math"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func newIntSeq() *multiWriterBtreeSeq {
	return newMultiWriterBtreeSeq(func(i interface{}) uint32 {
		return uint32(i.(int))
	})
}

func Test_treeSeq(t *testing.T) {
	seq := newIntSeq()

	go func() {

		r := rand.New(rand.NewSource(time.Now().Unix()))
		for _, v := range r.Perm(127) {
			v = v + 1
			if err := seq.Put(v); err != nil {
				t.Error("put", err)
			}
		}
		seq.Close()
	}()

	var c int
	for c = 1; true; c++ {
		v := seq.GetNext()
		if v == nil {
			break
		}
		if c != v.(int) {
			t.Errorf("sequence is not in order, expected=%d, received=%d", c, v.(int))
		}
	}

	if c != 128 {
		t.Errorf("sequence did not complete correctly, expected=%d, received=%d", 128, c)
	}
}

func Test_treeSeqSync(t *testing.T) {
	seq := newIntSeq()

	go func() {

		wg := sync.WaitGroup{}
		wg.Add(127)

		r := rand.New(rand.NewSource(time.Now().Unix()))
		for _, v := range r.Perm(127) {
			go func(i int) {
				i = i + 1
				if err := seq.Put(i); err != nil {
					t.Error(err)
				}
				wg.Done()
			}(v)
		}

		wg.Wait()
		seq.Close()
	}()

	c := 1
	for ; true; c++ {
		v := seq.GetNext()
		if v == nil {
			break
		}
		if c != v.(int) {
			t.Errorf("sequence is not in order, expected=%d, received=%d", c, v.(int))
		}
	}
	if c != 128 {
		t.Errorf("sequence did not complete correctly, expected=%d, received=%d", 128, c)
	}
}

func Test_treeSeqClosed(t *testing.T) {
	seq := newIntSeq()

	seq.Close()
	if err := seq.Put(128); err == nil {
		t.Error("error expected")
	}

	seq.Close()
	v := seq.GetNext()
	if v != nil {
		t.Error("not nil from closed sequencer")
	}
}

func Test_treeSeqPreloaded(t *testing.T) {
	seq := newIntSeq()

	const BufferThreshold = 5000

	go func() {

		r := rand.New(rand.NewSource(time.Now().Unix()))
		for _, v := range r.Perm(BufferThreshold - 1) {
			v = v + 1
			if err := seq.Put(v); err != nil {
				t.Error("put", err)
			}
		}
		seq.Close()

	}()

	var c int
	for c = 1; true; c++ {
		v := seq.GetNext()
		if v == nil {
			break
		}
		if c != v.(int) {
			t.Errorf("sequence is not in order, expected=%d, received=%d", c, v.(int))
		}
	}

	if c != BufferThreshold {
		t.Errorf("sequence did not complete correctly, expected=%d, received=%d", BufferThreshold, c)
	}
}

func newMultiWriterBtreeSeq(seqF func(interface{}) uint32) *multiWriterBtreeSeq {
	seq := &multiWriterBtreeSeq{
		singleWriterBtreeSeq: &singleWriterBtreeSeq{
			ch:            make(chan interface{}),
			tree:          btree.NewWith(4, utils.UInt32Comparator),
			nextSeq:       1,
			maxOutOfOrder: math.MaxUint32,
		},
		writeCh: make(chan *multiWriterSeqEntry),
	}

	go func() {
		for entry := range seq.writeCh {
			_ = seq.PutSequenced(seqF(entry.val), entry.val)
			close(entry.doneC)
		}
	}()

	return seq
}

// multiWriterBtreeSeq extends singleWriterBtreeSeq to be a multi-writer, multi-reader capable sequencer
type multiWriterBtreeSeq struct {
	*singleWriterBtreeSeq
	writeCh chan *multiWriterSeqEntry
}

type multiWriterSeqEntry struct {
	val   interface{}
	doneC chan struct{}
}

func (seq *multiWriterBtreeSeq) Put(val interface{}) error {
	if seq.closed.Get() {
		return errors.New("can't write to closed sequencer")
	}

	entry := &multiWriterSeqEntry{
		val:   val,
		doneC: make(chan struct{}),
	}
	seq.writeCh <- entry
	<-entry.doneC
	return nil
}
