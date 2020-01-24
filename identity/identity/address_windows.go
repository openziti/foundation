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
	"github.com/pkg/errors"
	"net/url"
	"strings"
)

func parseAddr(addr string) (*url.URL, error) {
	if !isFile(addr) {
		if strings.HasPrefix(addr, "pem:") {
			return &url.URL{
				Scheme:     "pem",
				Opaque:     strings.TrimPrefix(addr, "pem:"),
			}, nil
		}

		return nil, errors.New("non-file address supplied, but not supported [%s]")
	}

	if strings.HasPrefix(addr, "file://") {
		addr = strings.Replace(addr, "file://", "", 1)
	}
	return &url.URL{
		Scheme: "file",
		Path:   addr,
	}, nil
}

func isFile(addr string) bool {
	return !(isPem(addr) || isEngine(addr))
}

func isPem(addr string) bool {
	return strings.HasPrefix(addr, "pem:")
}

func isEngine(addr string) bool {
	return strings.HasPrefix(addr, "engine:")
}
