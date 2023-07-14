// Copyright  The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package prometheusremotewritereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusremotewritereceiver"

import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusremotewritereceiver/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const (
	defaultBindEndpoint  = "0.0.0.0:19291"
	defaultTimeThreshold = 24
)

// NewFactory creates a new Prometheus remote write receiver factory.
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		createDefaultConfig,
		receiver.WithMetrics(createMetricsReceiver, metadata.MetricsStability))
}

func createDefaultConfig() component.Config {
	return &Config{
		HTTPServerSettings: confighttp.HTTPServerSettings{
			Endpoint: defaultBindEndpoint,
		},
		TimeThreshold: defaultTimeThreshold,
	}
}

// createMetricsReceiver creates a metrics receiver based on provided config.
func createMetricsReceiver(
	_ context.Context,
	params receiver.CreateSettings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (receiver.Metrics, error) {
	if nextConsumer == nil {
		return nil, component.ErrNilNextConsumer
	}
	return newReceiver(params, cfg.(*Config), nextConsumer)
}
