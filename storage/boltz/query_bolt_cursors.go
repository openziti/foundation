/*
	Copyright 2019 Netfoundry, Inc.

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

package boltz

import "go.etcd.io/bbolt"

type BaseBoltCursor struct {
	cursor *bbolt.Cursor
	key    []byte
}

func (f *BaseBoltCursor) IsValid() bool {
	return f.key != nil
}

func (f *BaseBoltCursor) Id() []byte {
	return f.key
}

type ForwardBoltCursor struct {
	BaseBoltCursor
}

func (f *ForwardBoltCursor) Init() {
	f.key, _ = f.cursor.First()
}

func (f *ForwardBoltCursor) Next() {
	f.key, _ = f.cursor.Next()
}

type ReverseBoltCursor struct {
	BaseBoltCursor
}

func (f *ReverseBoltCursor) Init() {
	f.key, _ = f.cursor.Last()
}

func (f *ReverseBoltCursor) Next() {
	f.key, _ = f.cursor.Prev()
}

type ForwardIndexBoltCursor struct {
	BaseBoltCursor
}

func (f *ForwardIndexBoltCursor) Init() {
	f.key, _ = f.cursor.First()
}

func (f *ForwardIndexBoltCursor) Next() {
	_, f.key = f.cursor.Next()
}

type ReverseIndexBoltCursor struct {
	BaseBoltCursor
}

func (f *ReverseIndexBoltCursor) Init() {
	f.key, _ = f.cursor.Last()
}

func (f *ReverseIndexBoltCursor) Next() {
	_, f.key = f.cursor.Prev()
}
