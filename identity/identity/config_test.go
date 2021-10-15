package identity

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
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

const (
	TestValueCert       = "./ziti/etc/ca/intermediate/certs/ctrl-client.cert.pem"
	TestValueKey        = "./ziti/etc/ca/intermediate/private/ctrl.key.pem"
	TestValueServerCert = "./ziti/etc/ca/intermediate/certs/ctrl-server.cert.pem"
	TestValueServerKey  = "./ziti/etc/ca/intermediate/certs/ctrl-server.key.pem"
	TestValueCa         = "./ziti/etc/ca/intermediate/certs/ca-chain.cert.pem"

	TestValuePathContext = "my.path"

	TestValueMissingOrBlankFieldErrorTemplate = "required configuration value [%s] is missing or is blank"
	TestValueMissingOrBlankFieldsTemplate     = "required configuration values [%s], [%s] are both missing or are blank"
	TestValueMapStringErrorTemplate           = "value [%s] must be a string"

	TestValueJsonTemplate                     = `
		{
		  "cert": "%s",
		  "key": "%s",
		  "server_cert": "%s",
		  "server_key": "%s",
		  "ca": "%s"
		}`

	TestValueYamlTemplate = `
cert: "%s"
key: "%s"
server_cert: "%s"
server_key: "%s"
ca: "%s"
`
)

func Test_Config(t *testing.T) {
	t.Run("can parse from JSON", func(t *testing.T) {
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, TestValueServerCert, TestValueServerKey, TestValueCa))

		config := Config{}
		err := json.Unmarshal(identityConfigJson, &config)

		req := require.New(t)
		req.NoError(err)
		req.Equal(config.Cert, TestValueCert)
		req.Equal(config.Key, TestValueKey)
		req.Equal(config.ServerCert, TestValueServerCert)
		req.Equal(config.ServerKey, TestValueServerKey)
		req.Equal(config.CA, TestValueCa)
	})

	t.Run("can parse from YAML", func(t *testing.T) {
		identityConfigYaml := []byte(fmt.Sprintf(TestValueYamlTemplate, TestValueCert, TestValueKey, TestValueServerCert, TestValueServerKey, TestValueCa))

		config := Config{}
		err := yaml.Unmarshal(identityConfigYaml, &config)

		req := require.New(t)
		req.NoError(err)
		req.Equal(config.Cert, TestValueCert)
		req.Equal(config.Key, TestValueKey)
		req.Equal(config.ServerCert, TestValueServerCert)
		req.Equal(config.ServerKey, TestValueServerKey)
		req.Equal(config.CA, TestValueCa)
	})
}

func Test_Config_Validate(t *testing.T) {
	t.Run("all fields present returns no errors", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, TestValueServerCert, TestValueServerKey, TestValueCa))
		config := Config{}
		err := json.Unmarshal(identityConfigJson, &config)

		req.NoError(err)
		req.NoError(config.Validate())
	})

	t.Run("all fields present returns no errors", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, TestValueServerCert, TestValueServerKey, TestValueCa))
		config := Config{}
		err := json.Unmarshal(identityConfigJson, &config)

		req.NoError(err)
		req.NoError(config.Validate())
	})

	t.Run("empty string cert returns error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, "", TestValueKey, TestValueServerCert, TestValueServerKey, TestValueCa))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.Validate()
		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMissingOrBlankFieldErrorTemplate, "cert"))
	})

	t.Run("empty string key returns error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, "", TestValueServerCert, TestValueServerKey, TestValueCa))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.Validate()
		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMissingOrBlankFieldErrorTemplate, "key"))
	})

	t.Run("empty string server_cert returns error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, "", TestValueServerKey, TestValueCa))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.Validate()
		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMissingOrBlankFieldErrorTemplate, "server_cert"))
	})

	t.Run("empty string ca returns error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, TestValueServerCert, TestValueServerKey, ""))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.Validate()
		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMissingOrBlankFieldErrorTemplate, "ca"))
	})

	t.Run("empty string server_key returns no error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, TestValueServerCert, "", TestValueCa))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.Validate()
		req.NoError(err)
	})
}

func Test_Config_ValidateWithPathContext(t *testing.T) {
	t.Run("all fields present returns no errors", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, TestValueServerCert, TestValueServerKey, TestValueCa))
		config := Config{}
		err := json.Unmarshal(identityConfigJson, &config)

		req.NoError(err)
		req.NoError(config.ValidateWithPathContext(TestValuePathContext))
	})

	t.Run("empty string cert returns error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, "", TestValueKey, TestValueServerCert, TestValueServerKey, TestValueCa))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateWithPathContext(TestValuePathContext)
		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMissingOrBlankFieldErrorTemplate, TestValuePathContext+".cert"))
	})

	t.Run("empty string key returns error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, "", TestValueServerCert, TestValueServerKey, TestValueCa))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateWithPathContext(TestValuePathContext)
		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMissingOrBlankFieldErrorTemplate, TestValuePathContext+".key"))
	})

	t.Run("empty string server_cert returns error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, "", TestValueServerKey, TestValueCa))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateWithPathContext(TestValuePathContext)
		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMissingOrBlankFieldErrorTemplate, TestValuePathContext+".server_cert"))
	})

	t.Run("empty string ca returns error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, TestValueServerCert, TestValueServerKey, ""))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateWithPathContext(TestValuePathContext)
		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMissingOrBlankFieldErrorTemplate, TestValuePathContext+".ca"))
	})

	t.Run("empty string server_key returns no error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, TestValueServerCert, "", TestValueCa))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateWithPathContext(TestValuePathContext)
		req.NoError(err)
	})
}

func Test_Config_ValidateForClient(t *testing.T) {
	t.Run("all fields present returns no errors", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, TestValueServerCert, TestValueServerKey, TestValueCa))
		config := Config{}
		err := json.Unmarshal(identityConfigJson, &config)

		req.NoError(err)
		req.NoError(config.ValidateForClient())
	})

	t.Run("minimum fields present returns no errors", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, "", "", ""))
		config := Config{}
		err := json.Unmarshal(identityConfigJson, &config)

		req.NoError(err)
		req.NoError(config.ValidateForClient())
	})

	t.Run("empty string cert returns error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, "", TestValueKey, TestValueServerCert, TestValueServerKey, TestValueCa))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateForClient()
		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMissingOrBlankFieldErrorTemplate, "cert"))
	})

	t.Run("empty string key returns error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, "", TestValueServerCert, TestValueServerKey, TestValueCa))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateForClient()
		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMissingOrBlankFieldErrorTemplate, "key"))
	})

	t.Run("empty string server_cert returns no error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, "", TestValueServerKey, TestValueCa))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateForClient()
		req.NoError(err)
	})

	t.Run("empty string ca returns no error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, TestValueServerCert, TestValueServerKey, ""))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateForClient()
		req.NoError(err)
	})

	t.Run("empty string server_key returns no error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, TestValueServerCert, "", TestValueCa))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateForClient()
		req.NoError(err)
	})
}

func Test_Config_ValidateForClientWithPathContext(t *testing.T) {

	t.Run("all fields present returns no errors", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, TestValueServerCert, TestValueServerKey, TestValueCa))
		config := Config{}
		err := json.Unmarshal(identityConfigJson, &config)

		req.NoError(err)
		req.NoError(config.ValidateForClientWithPathContext(TestValuePathContext))
	})

	t.Run("minimum fields present returns no errors", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, "", "", ""))
		config := Config{}
		err := json.Unmarshal(identityConfigJson, &config)

		req.NoError(err)
		req.NoError(config.ValidateForClientWithPathContext(TestValuePathContext))
	})

	t.Run("empty string cert returns error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, "", TestValueKey, TestValueServerCert, TestValueServerKey, TestValueCa))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateForClientWithPathContext(TestValuePathContext)
		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMissingOrBlankFieldErrorTemplate, TestValuePathContext+".cert"))
	})

	t.Run("empty string key returns error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, "", TestValueServerCert, TestValueServerKey, TestValueCa))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateForClientWithPathContext(TestValuePathContext)
		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMissingOrBlankFieldErrorTemplate, TestValuePathContext+".key"))
	})

	t.Run("empty string server_cert returns no error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, "", TestValueServerKey, TestValueCa))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateForClientWithPathContext(TestValuePathContext)
		req.NoError(err)
	})

	t.Run("empty string ca returns no error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, TestValueServerCert, TestValueServerKey, ""))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateForClientWithPathContext(TestValuePathContext)
		req.NoError(err)
	})

	t.Run("empty string server_key returns no error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, TestValueServerCert, "", TestValueCa))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateForClientWithPathContext(TestValuePathContext)
		req.NoError(err)
	})
}

func Test_Config_ValidateForServer(t *testing.T) {
	t.Run("all fields present returns no errors", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, TestValueServerCert, TestValueServerKey, TestValueCa))
		config := Config{}
		err := json.Unmarshal(identityConfigJson, &config)

		req.NoError(err)
		req.NoError(config.ValidateForServer())
	})

	t.Run("minimum fields present returns no errors", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, "", "", TestValueServerCert, TestValueServerKey, TestValueCa))
		config := Config{}
		err := json.Unmarshal(identityConfigJson, &config)

		req.NoError(err)
		req.NoError(config.ValidateForServer())
	})

	t.Run("minimum fields present no server_key returns no errors", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, "", TestValueKey, TestValueServerCert, "", TestValueCa))
		config := Config{}
		err := json.Unmarshal(identityConfigJson, &config)

		req.NoError(err)
		req.NoError(config.ValidateForServer())
	})

	t.Run("empty string cert returns no error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, "", TestValueKey, TestValueServerCert, TestValueServerKey, TestValueCa))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateForServer()
		req.NoError(err)
	})

	t.Run("empty string key returns no error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, "", TestValueServerCert, TestValueServerKey, TestValueCa))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateForServer()
		req.NoError(err)
	})

	t.Run("empty string server_cert returns error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, "", TestValueServerKey, TestValueCa))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateForServer()
		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMissingOrBlankFieldErrorTemplate, "server_cert"))
	})

	t.Run("empty string ca returns error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, TestValueServerCert, TestValueServerKey, ""))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateForServer()
		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMissingOrBlankFieldErrorTemplate, "ca"))
	})

	t.Run("empty string server_key and no default key returns error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, "", TestValueServerCert, "", TestValueCa))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateForServer()
		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMissingOrBlankFieldsTemplate, "key", "server_key"))
	})
}

func Test_Config_ValidateForServerWithPathContext(t *testing.T) {
	t.Run("all fields present returns no errors", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, TestValueServerCert, TestValueServerKey, TestValueCa))
		config := Config{}
		err := json.Unmarshal(identityConfigJson, &config)

		req.NoError(err)
		req.NoError(config.ValidateForServerWithPathContext(TestValuePathContext))
	})

	t.Run("minimum fields present returns no errors", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, "", "", TestValueServerCert, TestValueServerKey, TestValueCa))
		config := Config{}
		err := json.Unmarshal(identityConfigJson, &config)

		req.NoError(err)
		req.NoError(config.ValidateForServerWithPathContext(TestValuePathContext))
	})

	t.Run("minimum fields present no server_key returns no errors", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, "", TestValueKey, TestValueServerCert, "", TestValueCa))
		config := Config{}
		err := json.Unmarshal(identityConfigJson, &config)

		req.NoError(err)
		req.NoError(config.ValidateForServerWithPathContext(TestValuePathContext))
	})

	t.Run("empty string cert returns no error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, "", TestValueKey, TestValueServerCert, TestValueServerKey, TestValueCa))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateForServerWithPathContext(TestValuePathContext)
		req.NoError(err)
	})

	t.Run("empty string key returns no error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, "", TestValueServerCert, TestValueServerKey, TestValueCa))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateForServerWithPathContext(TestValuePathContext)
		req.NoError(err)
	})

	t.Run("empty string server_cert returns error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, "", TestValueServerKey, TestValueCa))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateForServerWithPathContext(TestValuePathContext)
		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMissingOrBlankFieldErrorTemplate, TestValuePathContext+".server_cert"))
	})

	t.Run("empty string ca returns error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, TestValueKey, TestValueServerCert, TestValueServerKey, ""))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateForServerWithPathContext(TestValuePathContext)
		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMissingOrBlankFieldErrorTemplate, TestValuePathContext+".ca"))
	})

	t.Run("empty string server_key and no default key returns error", func(t *testing.T) {
		req := require.New(t)
		identityConfigJson := []byte(fmt.Sprintf(TestValueJsonTemplate, TestValueCert, "", TestValueServerCert, "", TestValueCa))
		config := Config{}

		err := json.Unmarshal(identityConfigJson, &config)
		req.NoError(err)

		err = config.ValidateForServerWithPathContext(TestValuePathContext)
		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMissingOrBlankFieldsTemplate, TestValuePathContext+".key", TestValuePathContext+".server_key"))
	})
}

func Test_NewConfigFromMap(t *testing.T) {
	t.Run("can parse all values from a map", func(t *testing.T) {
		req := require.New(t)

		configMap := map[interface{}]interface{}{
			"cert":        TestValueCert,
			"key":         TestValueKey,
			"server_cert": TestValueServerCert,
			"server_key":  TestValueServerKey,
			"ca":          TestValueCa,
		}

		config, err := NewConfigFromMap(configMap)

		req.NoError(err)
		req.Equal(config.Cert, TestValueCert)
		req.Equal(config.Key, TestValueKey)
		req.Equal(config.ServerCert, TestValueServerCert)
		req.Equal(config.ServerKey, TestValueServerKey)
		req.Equal(config.CA, TestValueCa)
	})

	t.Run("errors on non-string cert", func(t *testing.T) {
		req := require.New(t)

		configMap := map[interface{}]interface{}{
			"cert":        1,
			"key":         TestValueKey,
			"server_cert": TestValueServerCert,
			"server_key":  TestValueServerKey,
			"ca":          TestValueCa,
		}

		_, err := NewConfigFromMap(configMap)

		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMapStringErrorTemplate, "cert"))
	})

	t.Run("errors on non-string key", func(t *testing.T) {
		req := require.New(t)

		configMap := map[interface{}]interface{}{
			"cert":        TestValueCert,
			"key":         1,
			"server_cert": TestValueServerCert,
			"server_key":  TestValueServerKey,
			"ca":          TestValueCa,
		}

		_, err := NewConfigFromMap(configMap)

		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMapStringErrorTemplate, "key"))
	})

	t.Run("errors on non-string server_cert", func(t *testing.T) {
		req := require.New(t)

		configMap := map[interface{}]interface{}{
			"cert":        TestValueCert,
			"key":         TestValueKey,
			"server_cert": 1,
			"server_key":  TestValueServerKey,
			"ca":          TestValueCa,
		}

		_, err := NewConfigFromMap(configMap)

		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMapStringErrorTemplate, "server_cert"))
	})

	t.Run("errors on non-string server_key", func(t *testing.T) {
		req := require.New(t)

		configMap := map[interface{}]interface{}{
			"cert":        TestValueCert,
			"key":         TestValueKey,
			"server_cert": TestValueServerCert,
			"server_key":  1,
			"ca":          TestValueCa,
		}

		_, err := NewConfigFromMap(configMap)

		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMapStringErrorTemplate, "server_key"))
	})

	t.Run("errors on non-string ca", func(t *testing.T) {
		req := require.New(t)

		configMap := map[interface{}]interface{}{
			"cert":        TestValueCert,
			"key":         TestValueKey,
			"server_cert": TestValueServerCert,
			"server_key":  TestValueServerKey,
			"ca":          1,
		}

		_, err := NewConfigFromMap(configMap)

		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMapStringErrorTemplate, "ca"))
	})
}

func Test_NewConfigFromMapWithPathContext(t *testing.T) {
	t.Run("can parse all values from a map", func(t *testing.T) {
		req := require.New(t)

		configMap := map[interface{}]interface{}{
			"cert":        TestValueCert,
			"key":         TestValueKey,
			"server_cert": TestValueServerCert,
			"server_key":  TestValueServerKey,
			"ca":          TestValueCa,
		}

		config, err := NewConfigFromMapWithPathContext(configMap, TestValuePathContext)

		req.NoError(err)
		req.Equal(config.Cert, TestValueCert)
		req.Equal(config.Key, TestValueKey)
		req.Equal(config.ServerCert, TestValueServerCert)
		req.Equal(config.ServerKey, TestValueServerKey)
		req.Equal(config.CA, TestValueCa)
	})

	t.Run("errors on non-string cert", func(t *testing.T) {
		req := require.New(t)

		configMap := map[interface{}]interface{}{
			"cert":        1,
			"key":         TestValueKey,
			"server_cert": TestValueServerCert,
			"server_key":  TestValueServerKey,
			"ca":          TestValueCa,
		}

		_, err := NewConfigFromMapWithPathContext(configMap, TestValuePathContext)

		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMapStringErrorTemplate, TestValuePathContext + ".cert"))
	})

	t.Run("errors on non-string key", func(t *testing.T) {
		req := require.New(t)

		configMap := map[interface{}]interface{}{
			"cert":        TestValueCert,
			"key":         1,
			"server_cert": TestValueServerCert,
			"server_key":  TestValueServerKey,
			"ca":          TestValueCa,
		}

		_, err := NewConfigFromMapWithPathContext(configMap, TestValuePathContext)

		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMapStringErrorTemplate, TestValuePathContext + ".key"))
	})

	t.Run("errors on non-string server_cert", func(t *testing.T) {
		req := require.New(t)

		configMap := map[interface{}]interface{}{
			"cert":        TestValueCert,
			"key":         TestValueKey,
			"server_cert": 1,
			"server_key":  TestValueServerKey,
			"ca":          TestValueCa,
		}

		_, err := NewConfigFromMapWithPathContext(configMap, TestValuePathContext)

		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMapStringErrorTemplate, TestValuePathContext + ".server_cert"))
	})

	t.Run("errors on non-string server_key", func(t *testing.T) {
		req := require.New(t)

		configMap := map[interface{}]interface{}{
			"cert":        TestValueCert,
			"key":         TestValueKey,
			"server_cert": TestValueServerCert,
			"server_key":  1,
			"ca":          TestValueCa,
		}

		_, err := NewConfigFromMapWithPathContext(configMap, TestValuePathContext)

		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMapStringErrorTemplate, TestValuePathContext + ".server_key"))
	})

	t.Run("errors on non-string ca", func(t *testing.T) {
		req := require.New(t)

		configMap := map[interface{}]interface{}{
			"cert":        TestValueCert,
			"key":         TestValueKey,
			"server_cert": TestValueServerCert,
			"server_key":  TestValueServerKey,
			"ca":          1,
		}

		_, err := NewConfigFromMapWithPathContext(configMap, TestValuePathContext)

		req.Error(err)
		req.Equal(err.Error(), fmt.Sprintf(TestValueMapStringErrorTemplate, TestValuePathContext + ".ca"))
	})
}
