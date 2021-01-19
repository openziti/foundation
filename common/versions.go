package common

import (
	"fmt"
	"strings"
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

type VersionProvider interface {
	Version() string
	BuildDate() string
	Revision() string
	AsVersionInfo() *VersionInfo
	EncoderDecoder() VersionEncDec
}

type VersionEncDec interface {
	Encode(*VersionInfo) ([]byte, error)
	Decode([]byte) (*VersionInfo, error)
}

type VersionInfo struct {
	Version   string
	Revision  string
	BuildDate string
	OS        string
	Arch      string
}

type VersionEncDecImpl struct{}

var StdVersionEncDec = VersionEncDecImpl{}

func (encDec *VersionEncDecImpl) Encode(info *VersionInfo) ([]byte, error) {
	out := fmt.Sprintf("%v|%v|%v|%v|%v", info.Version, info.Revision, info.BuildDate, info.OS, info.Arch)
	return []byte(out), nil
}

func (encDec *VersionEncDecImpl) Decode(info []byte) (*VersionInfo, error) {
	values := strings.Split(string(info), "|")

	if len(values) != 5 {
		return nil, fmt.Errorf("could not parse version info, expected 5 values got %d", len(values))
	}

	return &VersionInfo{
		Version:   values[0],
		Revision:  values[1],
		BuildDate: values[2],
		OS:        values[3],
		Arch:      values[4],
	}, nil

}
