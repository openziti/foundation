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

package channel2

import "encoding/json"

type Options struct {
	OutQueueSize int
	BindHandlers []BindHandler
	PeekHandlers []PeekHandler
}

func DefaultOptions() *Options {
	return &Options{
		OutQueueSize: 4,
	}
}

func LoadOptions(data map[interface{}]interface{}) *Options {
	options := DefaultOptions()

	if value, found := data["outQueueSize"]; found {
		if floatValue, ok := value.(float64); ok {
			options.OutQueueSize = int(floatValue)
		}
	}

	return options
}

func (o Options) String() string {
	data, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(data)
}