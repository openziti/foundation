/*
	Copyright NetFoundry Inc.

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

package concurrenz

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_WaitGroupCloseFirst(t *testing.T) {
	wg := NewWaitGroup()

	for i := 0; i < 10; i++ {
		notifier := make(chan struct{})
		close(notifier)
		wg.AddNotifier(notifier)
	}

	assert.True(t, wg.WaitForDone(time.Second))
}

func Test_WaitGroupTimed(t *testing.T) {
	wg := NewWaitGroup()
	var notifiers []chan struct{}

	for i := 0; i < 10; i++ {
		notifier := make(chan struct{})
		notifiers = append(notifiers, notifier)
		wg.AddNotifier(notifier)
	}

	for i := 0; i < 10; i++ {
		idx := i
		go func() {
			d := time.Duration(idx+1) * 10 * time.Millisecond
			time.Sleep(d)
			close(notifiers[idx])
			fmt.Printf("%v: closed after %v\n", idx, d)
		}()
	}

	start := time.Now()
	assert.True(t, wg.WaitForDone(time.Second))
	elapsed := time.Since(start)
	fmt.Printf("elapsed: %v\n", elapsed)
	assert.True(t, elapsed >= 100*time.Millisecond)
	assert.True(t, elapsed <= 150*time.Millisecond)
}

func Test_WaitGroupTimedHalf(t *testing.T) {
	wg := NewWaitGroup()
	var notifiers []chan struct{}

	for i := 0; i < 10; i++ {
		notifier := make(chan struct{})
		notifiers = append(notifiers, notifier)
		wg.AddNotifier(notifier)

		if i%2 == 0 {
			close(notifier)
		}
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		idx := i
		go func() {
			d := time.Duration(idx+1) * 10 * time.Millisecond
			time.Sleep(d)
			close(notifiers[idx])
			fmt.Printf("%v: closed after %v\n", idx, d)
		}()
	}

	start := time.Now()
	assert.True(t, wg.WaitForDone(time.Second))
	elapsed := time.Since(start)
	fmt.Printf("elapsed: %v\n", elapsed)
	assert.True(t, elapsed >= 100*time.Millisecond)
	assert.True(t, elapsed <= 150*time.Millisecond)
}

func Test_WaitGroupTimout(t *testing.T) {
	wg := NewWaitGroup()
	var notifiers []chan struct{}

	for i := 0; i < 10; i++ {
		notifier := make(chan struct{})
		notifiers = append(notifiers, notifier)
		wg.AddNotifier(notifier)
	}

	for i := 0; i < 10; i++ {
		idx := i
		go func() {
			d := time.Duration(idx+1) * 10 * time.Millisecond
			time.Sleep(d)
			close(notifiers[idx])
			fmt.Printf("%v: closed after %v\n", idx, d)
		}()
	}

	start := time.Now()
	assert.False(t, wg.WaitForDone(50*time.Millisecond))
	elapsed := time.Since(start)
	fmt.Printf("elapsed: %v\n", elapsed)
	assert.True(t, elapsed >= 50*time.Millisecond)
	assert.True(t, elapsed <= 60*time.Millisecond)
}
