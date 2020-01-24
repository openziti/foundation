// +build windows

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

package identity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	filePrefix                   = "file://"
	fileScheme                   = "file"
	absolutePathUpperDriveLetter = "C:\\some\\dir"
	absolutePathLowerDriveLetter = "c:\\some\\dir"
	relativePath                 = "some\\dir"
)

func TestParseAddrWithRelativeWindowsPathNoFilePrefix(t *testing.T) {
	r, err := parseAddr(relativePath)

	assert.NoError(t, err)
	assert.Equal(t, fileScheme, r.Scheme)
	assert.Equal(t, relativePath, r.Path)
}

func TestParseAddrWithRelativeWindowsPathWithFilePrefix(t *testing.T) {
	r, err := parseAddr(filePrefix + relativePath)

	assert.NoError(t, err)
	assert.Equal(t, fileScheme, r.Scheme)
	assert.Equal(t, relativePath, r.Path)
}

func TestParseAddrWithAbsoluteWindowsPathWithNoFilePrefixWithUpperDriveLetter(t *testing.T) {
	r, err := parseAddr(absolutePathUpperDriveLetter)

	assert.NoError(t, err)
	assert.Equal(t, fileScheme, r.Scheme)
	assert.Equal(t, r.Path, absolutePathUpperDriveLetter)
}

func TestParseAddrWithAbsoluteWindowsPathWithFilePrefixWithUpperDriveLetter(t *testing.T) {
	r, err := parseAddr(filePrefix + absolutePathUpperDriveLetter)

	assert.NoError(t, err)
	assert.Equal(t, fileScheme, r.Scheme)
	assert.Equal(t, r.Path, absolutePathUpperDriveLetter)
}

func TestParseAddrWithAbsoluteWindowsPathWithNoFilePrefixWithLowerDriveLetter(t *testing.T) {
	r, err := parseAddr(absolutePathLowerDriveLetter)

	assert.NoError(t, err)
	assert.Equal(t, fileScheme, r.Scheme)
	assert.Equal(t, r.Path, absolutePathLowerDriveLetter)
}

func TestParseAddrWithAbsoluteWindowsPathWithFilePrefixWithLowerDriveLetter(t *testing.T) {
	r, err := parseAddr(filePrefix + absolutePathLowerDriveLetter)

	assert.NoError(t, err)
	assert.Equal(t, fileScheme, r.Scheme)
	assert.Equal(t, r.Path, absolutePathLowerDriveLetter)
}
