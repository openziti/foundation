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

package config

import (
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"reflect"
	"time"
)

func CheckFlag(options map[interface{}]interface{}, name string) bool {
	configVal, ok := options[name]
	if !ok {
		return true
	}
	result, ok := configVal.(bool)
	return !ok || result // True if value cannot be parsed or is set to true
}

func GetDuration(config map[interface{}]interface{}, name string, defaultVal time.Duration) time.Duration {
	intfVal, ok := config[name]
	if !ok {
		pfxlog.Logger().Debugf("No config value for %v specified, using default value of %v", name, defaultVal.String())
		return defaultVal
	}

	strVal, ok := intfVal.(string)

	if !ok {
		err := fmt.Errorf("config value %v must be a string but it is illegal type %v instead", name, reflect.TypeOf(intfVal).Name())
		panic(err)
	}

	result, err := time.ParseDuration(strVal)
	if err != nil {
		panic(fmt.Errorf("failed to parse duration from config value %v: %v", name, err))
	}
	return result
}