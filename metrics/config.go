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

package metrics

import (
	"errors"
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"time"
)

type Config struct {
	handlers       map[Handler]Handler
	ReportInterval time.Duration
}

func LoadConfig(srcmap map[interface{}]interface{}) (*Config, error) {
	cfg := &Config{
		handlers:       make(map[Handler]Handler),
		ReportInterval: 15 * time.Second,
	}

	pfxlog.Logger().Infof("Loading metrics configs")

	for k, v := range srcmap {
		if name, ok := k.(string); ok {
			switch name {

			// testing file output
			case string(HandlerTypeJSONFile):

				if submap, ok := v.(map[interface{}]interface{}); ok {
					if jsonfileConfig, err := LoadJSONFileConfig(submap); err == nil {
						if jsonFileHandler, err := NewJsonFileMetricsHandler(jsonfileConfig); err == nil {
							cfg.handlers[jsonFileHandler] = jsonFileHandler
							pfxlog.Logger().Infof("added JSON File handler")
						} else {
							return nil, fmt.Errorf("error creating JSON File handler (%s)", err)
						}
					} else {
						pfxlog.Logger().Warnf("Error loading the JSON File handler: (%s)", err)
					}
				} else {
					return nil, errors.New("invalid config for JSON File Handler ")
				}

			case string(HandlerTypeInfluxDB):
				if submap, ok := v.(map[interface{}]interface{}); ok {
					if influxCfg, err := LoadInfluxConfig(submap); err == nil {
						if influxHandler, err := NewInfluxDBMetricsHandler(influxCfg); err == nil {
							cfg.handlers[influxHandler] = influxHandler
							pfxlog.Logger().Infof("added influx handler")
						} else {
							return nil, fmt.Errorf("error creating influx handler (%s)", err)
						}
					}
				} else {
					return nil, errors.New("invalid influx stanza")
				}
			case "reportInterval":
				val, ok := v.(string)
				if !ok {
					return nil, errors.New("metrics.reportInterval must be a string duration, for example: 15s")
				}
				interval, err := time.ParseDuration(val)
				if err != nil {
					return nil, err
				}
				cfg.ReportInterval = interval
			}
		}
	}

	return cfg, nil
}
