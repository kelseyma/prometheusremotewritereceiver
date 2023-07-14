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
	"errors"
	"io"
	"net/http"
	"sync"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite"
	"github.com/prometheus/prometheus/storage/remote"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/obsreport"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"
)

const (
	receiverFormat = "protobuf"
)

type prometheusRemoteWriteReceiver struct {
	params       receiver.CreateSettings
	host         component.Host
	nextConsumer consumer.Metrics

	shutdownWG sync.WaitGroup
	cancelFunc context.CancelFunc

	config  *Config
	logger  *zap.Logger
	server  *http.Server
	obsrecv *obsreport.Receiver
}

func newReceiver(params receiver.CreateSettings, cfg *Config, consumer consumer.Metrics) (*prometheusRemoteWriteReceiver, error) {
	obsrecv, err := obsreport.NewReceiver(obsreport.ReceiverSettings{
		ReceiverID:             params.ID,
		ReceiverCreateSettings: params,
	})
	prwr := &prometheusRemoteWriteReceiver{
		params:       params,
		nextConsumer: consumer,
		config:       cfg,
		logger:       params.Logger,
		obsrecv:      obsrecv,
	}
	return prwr, err
}

// TODO: needs refinement (should contents be wrapped in startOnce?)
func (r *prometheusRemoteWriteReceiver) Start(_ context.Context, host component.Host) error {
	r.host = host
	_, cancel := context.WithCancel(context.Background())
	r.cancelFunc = cancel

	server, err := r.config.HTTPServerSettings.ToServer(r.host, r.params.TelemetrySettings, r,
		// TODO: copied from soc
		confighttp.WithDecoder("snappy", func(body io.ReadCloser) (io.ReadCloser, error) {
			return body, nil
		}))
	if err != nil {
		return err
	}
	r.server = server

	listener, err := r.config.HTTPServerSettings.ToListener()
	if err != nil {
		return err
	}
	r.shutdownWG.Add(1)
	go func() {
		defer r.shutdownWG.Done()
		if errHTTP := r.server.Serve(listener); errHTTP != nil && !errors.Is(errHTTP, http.ErrServerClosed) {
			host.ReportFatalError(errHTTP)
		}
	}()
	r.logger.Info("Started Prometheus Remote Write Receiver")

	return nil
}

func (r *prometheusRemoteWriteReceiver) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	ctx := r.obsrecv.StartMetricsOp(request.Context())
	req, err := remote.DecodeWriteRequest(request.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TODO: remove dev testing log
	r.logger.Debug("Received request:", zap.Any("request_body", req))

	if len(req.Timeseries) == 0 && len(req.Metadata) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	md, err := prometheusremotewrite.FromTimeSeries(req.Timeseries, prometheusremotewrite.Settings{
		TimeThreshold: r.config.TimeThreshold,
		Logger:        r.logger,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TODO: remove dev testing log
	r.logger.Debug("Processed metrics:", zap.Any("pmetric_metrics", md))

	metricCount := md.ResourceMetrics().Len()
	dataPointCount := md.DataPointCount()
	if metricCount != 0 {
		err = r.nextConsumer.ConsumeMetrics(ctx, md)
	}
	r.obsrecv.EndMetricsOp(ctx, receiverFormat, dataPointCount, err)
	w.WriteHeader(http.StatusAccepted)
}

// TODO: needs refinement (should contents be wrapped in stopOnce?)
func (r *prometheusRemoteWriteReceiver) Shutdown(ctx context.Context) error {
	if r.cancelFunc != nil {
		defer r.cancelFunc()
	}

	var err error
	if r.server != nil {
		err = r.server.Close()
	}
	r.shutdownWG.Wait()

	return err
}
