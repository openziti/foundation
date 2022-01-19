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

package channel

import (
	"errors"
	"testing"
)

func Test_getRetryVersionFor(t *testing.T) {
	twoAndOne := []uint32{2, 1}

	tests := []struct {
		name          string
		err           error
		localVersions []uint32
		want          uint32
		want1         bool
	}{
		struct {
			name          string
			err           error
			localVersions []uint32
			want          uint32
			want1         bool
		}{name: "non version error", err: errors.New("foo"), localVersions: twoAndOne, want: 1, want1: false},
		{name: "empty non version error", err: UnsupportedVersionError{}, localVersions: twoAndOne, want: 1, want1: false},
		{name: "v1", err: UnsupportedVersionError{supportedVersions: []uint32{1}}, localVersions: twoAndOne, want: 1, want1: true},
		{name: "v1, v2", err: UnsupportedVersionError{supportedVersions: []uint32{1, 2}}, localVersions: twoAndOne, want: 2, want1: true},
		{name: "v2, v1", err: UnsupportedVersionError{supportedVersions: []uint32{2, 1}}, localVersions: twoAndOne, want: 2, want1: true},
		{name: "v2", err: UnsupportedVersionError{supportedVersions: []uint32{2}}, localVersions: twoAndOne, want: 2, want1: true},
		{name: "v3", err: UnsupportedVersionError{supportedVersions: []uint32{3}}, localVersions: twoAndOne, want: 1, want1: false},
		{name: "v1, v2, v3", err: UnsupportedVersionError{supportedVersions: []uint32{1, 2, 3}}, localVersions: twoAndOne, want: 2, want1: true},
		{name: "v3, v2, v1", err: UnsupportedVersionError{supportedVersions: []uint32{3, 2, 1}}, localVersions: twoAndOne, want: 2, want1: true},
		{name: "v3, v1", err: UnsupportedVersionError{supportedVersions: []uint32{1, 3}}, localVersions: twoAndOne, want: 1, want1: true},
		{name: "v1, v3", err: UnsupportedVersionError{supportedVersions: []uint32{3, 1}}, localVersions: twoAndOne, want: 1, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getRetryVersionFor(tt.err, 1, tt.localVersions...)
			if got != tt.want {
				t.Errorf("getRetryVersionFor() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getRetryVersionFor() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
