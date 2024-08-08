package versions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

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

func Test_Compare(t *testing.T) {
	req := require.New(t)
	v1, err := ParseSemVer("v0.18.4")
	req.NoError(err)
	req.Equal("0.18.4", v1.String())
	req.True(v1.Equals(MustParseSemVer("0.18.4")))

	v2, err := ParseSemVer("0.18.5")
	req.NoError(err)
	req.Equal("0.18.5", v2.String())
	req.True(v2.Equals(MustParseSemVer("0.18.5")))

	v3, err := ParseSemVer("0.18.6")
	req.NoError(err)
	req.Equal("0.18.6", v3.String())
	req.True(v3.Equals(MustParseSemVer("0.18.6")))

	v4, err := ParseSemVer("0.19.6")
	req.NoError(err)
	req.Equal("0.19.6", v4.String())
	req.True(v4.Equals(MustParseSemVer("0.19.6")))

	v5, err := ParseSemVer("2.0.0")
	req.NoError(err)
	req.Equal("2.0.0", v5.String())
	req.True(v5.Equals(MustParseSemVer("2.0.0")))

	req.Equal(1, v2.CompareTo(v1))
	req.Equal(0, v2.CompareTo(v2))
	req.Equal(-1, v2.CompareTo(v3))
	req.Equal(-1, v2.CompareTo(v4))
	req.Equal(-1, v2.CompareTo(v5))

	versionInfo := &VersionInfo{Version: "v0.14.7"}
	req.False(versionInfo.HasMinimumVersion("0.18.5"))

	versionInfo = &VersionInfo{Version: "v0.18.4"}
	req.False(versionInfo.HasMinimumVersion("0.18.5"))

	versionInfo = &VersionInfo{Version: "v0.18.5"}
	req.True(versionInfo.HasMinimumVersion("0.18.5"))

	versionInfo = &VersionInfo{Version: "v0.19.0"}
	req.True(versionInfo.HasMinimumVersion("0.18.5"))

	versionInfo = &VersionInfo{Version: "v1.0.0"}
	req.True(versionInfo.HasMinimumVersion("0.18.5"))

	versionInfo = &VersionInfo{Version: "0.20"}
	req.True(versionInfo.HasMinimumVersion("0.18.5"))

	versionInfo = &VersionInfo{Version: "20"}
	req.True(versionInfo.HasMinimumVersion("0.18.5"))

	versionInfo = &VersionInfo{Version: "v0.0.0"}
	req.True(versionInfo.HasMinimumVersion("0.18.5"))

	// ensure invalid versions are still invalid
	_, err = ParseSemVer("2.0.0~alpha1")
	req.EqualError(err, `strconv.ParseInt: parsing "0~alpha1": invalid syntax`)

	_, err = ParseSemVer("2.0.0-rc2")
	req.EqualError(err, `strconv.ParseInt: parsing "0-rc2": invalid syntax`)
}
