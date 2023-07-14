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

package prometheusremotewrite // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite"

import (
	"fmt"
	"math"
	"strings"
	"time"

	"errors"

	"github.com/prometheus/prometheus/model/value"
	"github.com/prometheus/prometheus/prompb"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

func FromTimeSeries(timeseries []prompb.TimeSeries, settings Settings) (pmetric.Metrics, error) {
	md := pmetric.NewMetrics()
	for _, ts := range timeseries {
		empty := md.ResourceMetrics().AppendEmpty().ScopeMetrics().AppendEmpty().Metrics().AppendEmpty()
		newMetric := pmetric.NewMetric()

		// 1. get metric name from label
		metricName, err := getMetricNameLabel(ts.Labels)
		if err != nil {
			// TODO: from #14752 but should we just continue to next?
			return md, err
		}
		newMetric.SetName(metricName)

		// 2. determine metric type and unit
		// Diff from #14752: Split string instead of match regex so short metric names are also processed (e.g. redis_up, up), should verify/refine
		nameSplit := strings.Split(metricName, "_")
		mType := prompb.MetricMetadata_GAUGE
		if len(nameSplit) > 1 {
			lastSuffixInMetricName := nameSplit[len(nameSplit)-1]
			// TODO: logic from soc
			switch {
			case IsValidCumulativeSuffix(lastSuffixInMetricName):
				mType = prompb.MetricMetadata_COUNTER
			case lastSuffixInMetricName == "_info":
				mType = prompb.MetricMetadata_INFO
			}

			if IsValidUnit(lastSuffixInMetricName) {
				newMetric.SetUnit(lastSuffixInMetricName)
			} else if len(nameSplit) > 2 && IsValidUnit(nameSplit[len(nameSplit)-2]) {
				newMetric.SetUnit(nameSplit[len(nameSplit)-2])
			}
		}
		// TODO: remove dev testing log
		settings.Logger.Debug("Current metric:", zap.String("metric_name", newMetric.Name()), zap.String("metric_unit", newMetric.Unit()))

		// 3. add to pmetrics
		for _, s := range ts.Samples {
			point := pmetric.NewNumberDataPoint()
			if value.IsStaleNaN(s.Value) || math.IsNaN(s.Value) {
				continue
			} else {
				point.SetDoubleValue(s.Value)
			}

			point.SetTimestamp(pcommon.Timestamp(s.Timestamp * int64(time.Millisecond)))
			if point.Timestamp().AsTime().Before(time.Now().Add(-time.Duration(settings.TimeThreshold) * time.Hour)) {
				settings.Logger.Debug("Metric older than the threshold", zap.String("metric name", newMetric.Name()), zap.Time("metric_timestamp", point.Timestamp().AsTime()))
				continue
			}

			for _, l := range ts.Labels {
				if l.Name == nameStr || l.Value == "" {
					continue
				}
				point.Attributes().PutStr(l.Name, l.Value)
			}

			switch mType {
			case prompb.MetricMetadata_GAUGE:
				newMetric.SetEmptyGauge()
				point.CopyTo(newMetric.Gauge().DataPoints().AppendEmpty())
			case prompb.MetricMetadata_COUNTER:
				newMetric.SetEmptySum()
				newMetric.Sum().SetIsMonotonic(true)
				newMetric.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
				point.CopyTo(newMetric.Sum().DataPoints().AppendEmpty())
			}

			// TODO: remove dev testing log
			settings.Logger.Debug("Metric sample:",
				zap.String("metric_name", newMetric.Name()),
				zap.String("metric_unit", newMetric.Unit()),
				zap.Float64("metric_value", point.DoubleValue()),
				zap.Time("metric_timestamp", point.Timestamp().AsTime()),
				zap.String("metric_labels", fmt.Sprintf("%#v", point.Attributes())),
			)
		}
		newMetric.MoveTo(empty)
	}
	return md, nil
}

func getMetricNameLabel(labels []prompb.Label) (ret string, err error) {
	for _, label := range labels {
		if label.Name == nameStr {
			return label.Value, nil
		}
	}
	return "", errors.New("`__name__` label not found")
}

// IsValidCumulativeSuffix - TODO: update comment
func IsValidCumulativeSuffix(suffix string) bool {
	switch suffix {
	case
		"sum",
		"count",
		"counter",
		"total":
		return true
	}
	return false
}

// IsValidUnit - TODO: update comment
func IsValidUnit(unit string) bool {
	switch unit {
	case
		"seconds",
		"bytes":
		return true
	}
	return false
}
