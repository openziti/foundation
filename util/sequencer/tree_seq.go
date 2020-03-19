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
	"github.com/netfoundry/ziti-foundation/util/concurrenz"
	"github.com/emirpasic/gods/trees/btree"
	"github.com/emirpasic/gods/utils"
	"github.com/pkg/errors"
	"time"
)

func NewSingleWriterSeq(maxOutOfOrder uint32) Sequencer {
	return &singleWriterBtreeSeq{
		maxOutOfOrder: int(maxOutOfOrder),
		ch:            make(chan interface{}, 16),
		tree:          btree.NewWith(4, utils.UInt32Comparator),
		nextSeq:       1,
	}
}

// singleWriterBtreeSeq is a single write, multi reader capable sequencer
type singleWriterBtreeSeq struct {
	maxOutOfOrder int
	ch            chan interface{}
	tree          *btree.Tree
	nextSeq       uint32
	closed        concurrenz.AtomicBoolean
}

func (seq *singleWriterBtreeSeq) PutSequenced(itemSeq uint32, val interface{}) error {
	if seq.closed.Get() {
		return errors.New("can't write to closed sequencer")
	}
	if seq.nextSeq == itemSeq {
		seq.ch <- val
		seq.nextSeq++
		for !seq.tree.Empty() {
			nextKey := seq.tree.LeftKey().(uint32)
			if seq.nextSeq != nextKey {
				return nil
			}
			nextVal := seq.tree.LeftValue()
			seq.tree.Remove(nextKey)
			seq.ch <- nextVal
			seq.nextSeq++
		}
	} else if seq.tree.Size() < seq.maxOutOfOrder {
		seq.tree.Put(itemSeq, val)
	} else {
		return errors.Errorf("exceeded max out of order entries: %v", seq.maxOutOfOrder)
	}
	return nil
}

func (seq *singleWriterBtreeSeq) GetNext() interface{} {
	if val, ok := <-seq.ch; ok {
		return val
	}
	return nil
}

func (seq *singleWriterBtreeSeq) GetNextWithDeadline(t time.Time) (interface{}, error) {
	if t.IsZero() {
		result := seq.GetNext()
		if result == nil {
			return nil, ErrClosed
		}
		return result, nil
	}

	select {
	case v, open := <-seq.ch:
		if open {
			return v, nil
		}
		return nil, ErrClosed
	case <-time.After(time.Until(t)):
		return nil, ErrTimedOut
	}
}

func (seq *singleWriterBtreeSeq) Close() {
	if seq.closed.CompareAndSwap(false, true) {
		time.AfterFunc(time.Second, func() {
			close(seq.ch)
		})
	}
}