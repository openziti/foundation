/*
	Copyright 2019 NetFoundry, Inc.

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

import "testing"

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
