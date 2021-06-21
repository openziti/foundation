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
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/foundation/config"
	"time"
)

type Config struct {
	handlers       map[Handler]Handler
	Source         string
	Tags           map[string]string
	ReportInterval time.Duration
	EventSink      Handler
}

func LoadConfig(metricsConfig config.Map) *Config {
	cfg := &Config{
		handlers:       make(map[Handler]Handler),
		ReportInterval: metricsConfig.Duration("reportInterval", 15*time.Second),
	}

	pfxlog.Logger().Infof("Loading metrics configs")

	fileConfig := metricsConfig.Child(string(HandlerTypeJSONFile))
	if fileConfig == nil {
		fileConfig = metricsConfig.Child(string(HandlerTypeFile))
	}

	if fileConfig != nil {
		outputFileConfig := LoadOutputFileConfig(fileConfig)
		if !fileConfig.HasError() {
			outputFileHandler := NewOutputFileMetricsHandler(outputFileConfig)
			cfg.handlers[outputFileHandler] = outputFileHandler
			pfxlog.Logger().Infof("added metrics output file handler")
		}
	}

	if childMap := metricsConfig.Child(string(HandlerTypeInfluxDB)); childMap != nil {
		influxCfg := &influxConfig{
			url:      *childMap.RequireUrl("url"),
			database: childMap.RequireString("database"),
			username: childMap.RequireString("username"),
			password: childMap.RequireString("password"),
		}

		if !childMap.HasError() {
			if influxHandler, err := NewInfluxDBMetricsHandler(influxCfg); err == nil {
				cfg.handlers[influxHandler] = influxHandler
				pfxlog.Logger().Infof("added influx handler")
			} else {
				metricsConfig.SetError(err)
			}
		}
	}

	return cfg
}
