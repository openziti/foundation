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

package sequence

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestSequence(t *testing.T) {
	s := NewSequence()

	s1 := s.Next()
	if s1 != 1 {
		t.Errorf("expected [1], got [%d]", s1)
		return
	}

	h0, err := s.NextHash()
	if err != nil {
		t.Errorf("unexpected error (%s)", err)
		return
	}
	if len(h0) != 4 {
		t.Errorf("expected 4-character hash, got [%s]", h0)
		return
	}
}

func TestDups(t *testing.T) {
	s := NewSequence()

	results := make(chan map[string]struct{}, 4)
	errs := make(chan error, 4)

	for i := 0; i < 4; i++ {
		idx := i
		go func() {
			result := map[string]struct{}{}
			for j := 0; j < 100_000; j++ {
				next, err := s.NextHash()
				if err != nil {
					errs <- err
					return
				}
				if _, found := result[next]; found {
					errs <- errors.Errorf("id %v created twice", next)
					return
				}
				result[next] = struct{}{}
			}
			errs <- nil
			results <- result
			fmt.Printf("%v: key generation done\n", idx)
		}()
	}

	for i := 0; i < 4; i++ {
		err := <-errs
		if err != nil {
			t.Error(err)
		}
	}

	result := map[string]struct{}{}
	for i := 0; i < 4; i++ {
		next := <-results
		for id := range next {
			if _, found := result[id]; found {
				t.Error(errors.Errorf("id %v created twice", next))
			}
			result[id] = struct{}{}
		}
		fmt.Printf("%v: no dups found\n", i)
	}
}
