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

package channel2

import (
	"encoding/binary"
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

func TestMessageHeader_PutUint64Header(t *testing.T) {
	testHeader := int32(1111)
	testLEHeader := int32(2222)

	type args struct {
		key   int32
		value uint64
	}
	tests := []struct {
		name   string
		len    int
		args   args
	}{
		{name: "0 byte", len: 0, args: args{ key:testHeader, value: uint64(0)}},
		{name: "1 byte", len: 1, args: args{ key:testHeader, value: uint64(0x1)}},
		{name: "2 byte", len: 2, args: args{ key:testHeader, value: uint64(0x0201)}},
		{name: "3 byte", len: 3, args: args{ key:testHeader, value: uint64(0x030201)}},
		{name: "4 byte", len: 4, args: args{ key:testHeader, value: uint64(0x04030201)}},
		{name: "5 byte", len: 5, args: args{ key:testHeader, value: uint64(0x0504030201)}},
		{name: "6 byte", len: 6, args: args{ key:testHeader, value: uint64(0x060504030201)}},
		{name: "7 byte", len: 7, args: args{ key:testHeader, value: uint64(0x07060504000001)}},
		{name: "8 byte", len: 8, args: args{ key:testHeader, value: uint64(0x0807060504000000)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header := &MessageHeader{
				Headers:     make(map[int32][]byte),
			}

			header.PutUint64Header(tt.args.key, tt.args.value)
			leVal := make([]byte,8)
			binary.LittleEndian.PutUint64(leVal, tt.args.value)
			header.Headers[testLEHeader] = leVal

			binary, err := marshalHeaders(header.Headers)
			if err != nil {
				t.Error(tt.name, err)
			}

			header.Headers, err = unmarshalHeaders(binary)
			if err != nil {
				t.Error(tt.name, err)
			}

			if tt.len != len(header.Headers[testHeader]) {
				t.Errorf("%s: unexpected header len", tt.name)
			}

			v, found := header.GetUint64Header(tt.args.key)
			if !found {
				t.Errorf("%s: put failed", tt.name)
			}
			if v != tt.args.value {
				t.Errorf("%s: %d(%v) != %v", tt.name, v, header.Headers[tt.args.key], tt.args.value)
			}

			leV, found := header.GetUint64Header(testLEHeader)
			if !found {
				t.Errorf("%s: put failed", tt.name)
			}
			if leV != tt.args.value {
				t.Errorf("%s: %d(%v) != %v", tt.name, leV, header.Headers[tt.args.key], tt.args.value)
			}
		})
	}
}

func TestMessageHeader_PutUint32Header(t *testing.T) {
	testHeader := int32(1111)
	type args struct {
		key   int32
		value uint32
	}
	tests := []struct {
		name   string
		len    int
		args   args
	}{
		{name: "0 byte", len:0, args: args{ key:testHeader, value: uint32(0)}},
		{name: "1 byte", len:1, args: args{ key:testHeader, value: uint32(0x1)}},
		{name: "2 byte", len:2, args: args{ key:testHeader, value: uint32(0x0201)}},
		{name: "3 byte", len:3, args: args{ key:testHeader, value: uint32(0x030201)}},
		{name: "4 byte", len:4, args: args{ key:testHeader, value: uint32(0x04030201)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header := &MessageHeader{
				Headers:     make(map[int32][]byte),
			}

			header.PutUint32Header(tt.args.key, tt.args.value)
			if tt.len != len(header.Headers[testHeader]) {
				t.Errorf("%s: unexpected header len", tt.name)
			}
			v, found := header.GetUint32Header(tt.args.key)
			if !found {
				t.Errorf("%s: put failed", tt.name)
			}
			if v != tt.args.value {
				t.Errorf("%s: %d(%v) != %v", tt.name, v, header.Headers[tt.args.key], tt.args.value)
			}
		})
	}
}

func TestMessageHeader_PutUint16Header(t *testing.T) {
	testHeader := int32(1111)
	type args struct {
		key   int32
		value uint16
	}
	tests := []struct {
		name   string
		args   args
	}{
		{name: "0 byte", args: args{ key:testHeader, value: uint16(0)}},
		{name: "1 byte", args: args{ key:testHeader, value: uint16(0x1)}},
		{name: "2 byte", args: args{ key:testHeader, value: uint16(0xFF00)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header := &MessageHeader{
				Headers:     make(map[int32][]byte),
			}

			header.PutUint16Header(tt.args.key, tt.args.value)
			v, found := header.GetUint16Header(tt.args.key)
			if !found {
				t.Errorf("%s: put failed", tt.name)
			}
			if v != tt.args.value {
				t.Errorf("%s: %d(%v) != %v", tt.name, v, header.Headers[tt.args.key], tt.args.value)
			}
		})
	}
}