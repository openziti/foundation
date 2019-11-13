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

package dotziti

import (
	"github.com/netfoundry/ziti-foundation/identity/identity"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

func LoadIdentity(name string) (string, *identity.TokenId, error) {
	home := os.Getenv("ZITI_HOME")

	var err error
	idPath := ""
	if home != "" {
		idPath = filepath.Join(home, "identities.yml")
	} else {
		home, err = os.UserHomeDir()
		if err != nil {
			return "", nil, err
		}
		idPath = filepath.Join(home, ".ziti/identities.yml")
	}

	_, err = os.Stat(idPath)
	if err != nil {
		return "", nil, err
	}

	data, err := ioutil.ReadFile(idPath)
	if err != nil {
		return "", nil, err
	}

	var identities map[interface{}]interface{}
	err = yaml.Unmarshal(data, &identities)
	if err != nil {
		return "", nil, err
	}

	if value, found := identities[name]; found {
		idMap, ok := value.(map[interface{}]interface{})
		if !ok {
			return "", nil, errors.New("malformed identity")
		}

		endpointValue, found := idMap["endpoint"]
		if !found {
			return "", nil, errors.New("missing 'endpoint' in identity")
		}
		endpoint, ok := endpointValue.(string)
		if !ok {
			return "", nil, errors.New("invalid 'endpoint' in identity")
		}

		keyValue, found := idMap["key"]
		if !found {
			return "", nil, errors.New("missing 'key' in identity")
		}
		key, ok := keyValue.(string)
		if !ok {
			return "", nil, errors.New("invalid 'key' in identity")
		}

		certValue, found := idMap["cert"]
		if !found {
			return "", nil, errors.New("missing 'cert' in identity")
		}
		cert, ok := certValue.(string)
		if !ok {
			return "", nil, errors.New("invalid 'cert' in identity")
		}

		caCertValue, found := idMap["caCert"]
		if !found {
			return "", nil, errors.New("missing 'caCert' in identity")
		}
		caCert, ok := caCertValue.(string)
		if !ok {
			return "", nil, errors.New("invalid 'caCert' in identity")
		}

		if serverCertValue, found := idMap["serverCert"]; found {
			serverCert, ok := serverCertValue.(string)
			if !ok {
				return "", nil, errors.New("invalid 'serverCert' in identity")
			}
			id, err := identity.LoadServerIdentity(cert, serverCert, key, caCert)
			if err != nil {
				return "", nil, err
			}
			return endpoint, id, nil

		} else {
			id, err := identity.LoadClientIdentity(cert, key, caCert)
			if err != nil {
				return "", nil, err
			}
			return endpoint, id, nil
		}
	} else {
		return "", nil, fmt.Errorf("identity [%s] not found", name)
	}
}
