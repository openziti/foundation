package common

import (
	"github.com/stretchr/testify/require"
	"testing"
)

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

func Test_Versions_EncDec(t *testing.T) {

	const (
		version   = "v0.17.4"
		revision  = "0ec8cd4f0ccd"
		buildDate = "2020-11-13 00:47:59"
		os        = "linux"
		arch      = "amd64"
	)
	info := &VersionInfo{
		Version:   version,
		Revision:  revision,
		BuildDate: buildDate,
		OS:        os,
		Arch:      arch,
	}

	t.Run("can encode values", func(t *testing.T) {
		req := require.New(t)

		encodedVal, err := StdVersionEncDec.Encode(info)

		req.NoError(err)
		req.NotEmpty(encodedVal)

		t.Run("can decode values", func(t *testing.T) {
			req := require.New(t)

			outVal, err := StdVersionEncDec.Decode(encodedVal)

			req.NoError(err)
			req.NotNil(outVal)

			t.Run("version output matches input", func(t *testing.T) {
				req := require.New(t)
				req.Equal(version, outVal.Version)
			})

			t.Run("revision output matches input", func(t *testing.T) {
				req := require.New(t)
				req.Equal(revision, outVal.Revision)
			})

			t.Run("build date output matches input", func(t *testing.T) {
				req := require.New(t)
				req.Equal(buildDate, outVal.BuildDate)
			})

			t.Run("os output matches input", func(t *testing.T) {
				req := require.New(t)
				req.Equal(os, outVal.OS)
			})

			t.Run("arch output matches input", func(t *testing.T) {
				req := require.New(t)
				req.Equal(arch, outVal.Arch)
			})
		})
	})

}
