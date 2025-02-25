// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	conventions "go.opentelemetry.io/collector/semconv/v1.18.0"
)

type metricK8sNodeAllocatableCPU struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.node.allocatable_cpu metric with initial data.
func (m *metricK8sNodeAllocatableCPU) init() {
	m.data.SetName("k8s.node.allocatable_cpu")
	m.data.SetDescription("How many CPU cores remaining that the node can allocate to pods")
	m.data.SetUnit("{cores}")
	m.data.SetEmptyGauge()
}

func (m *metricK8sNodeAllocatableCPU) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sNodeAllocatableCPU) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sNodeAllocatableCPU) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sNodeAllocatableCPU(cfg MetricConfig) metricK8sNodeAllocatableCPU {
	m := metricK8sNodeAllocatableCPU{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricK8sNodeAllocatableEphemeralStorage struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.node.allocatable_ephemeral_storage metric with initial data.
func (m *metricK8sNodeAllocatableEphemeralStorage) init() {
	m.data.SetName("k8s.node.allocatable_ephemeral_storage")
	m.data.SetDescription("How many bytes of ephemeral storage remaining that the node can allocate to pods")
	m.data.SetUnit("By")
	m.data.SetEmptyGauge()
}

func (m *metricK8sNodeAllocatableEphemeralStorage) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sNodeAllocatableEphemeralStorage) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sNodeAllocatableEphemeralStorage) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sNodeAllocatableEphemeralStorage(cfg MetricConfig) metricK8sNodeAllocatableEphemeralStorage {
	m := metricK8sNodeAllocatableEphemeralStorage{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricK8sNodeAllocatableMemory struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.node.allocatable_memory metric with initial data.
func (m *metricK8sNodeAllocatableMemory) init() {
	m.data.SetName("k8s.node.allocatable_memory")
	m.data.SetDescription("How many bytes of RAM memory remaining that the node can allocate to pods")
	m.data.SetUnit("By")
	m.data.SetEmptyGauge()
}

func (m *metricK8sNodeAllocatableMemory) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sNodeAllocatableMemory) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sNodeAllocatableMemory) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sNodeAllocatableMemory(cfg MetricConfig) metricK8sNodeAllocatableMemory {
	m := metricK8sNodeAllocatableMemory{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricK8sNodeAllocatablePods struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.node.allocatable_pods metric with initial data.
func (m *metricK8sNodeAllocatablePods) init() {
	m.data.SetName("k8s.node.allocatable_pods")
	m.data.SetDescription("How many pods remaining the node can allocate")
	m.data.SetUnit("{pods}")
	m.data.SetEmptyGauge()
}

func (m *metricK8sNodeAllocatablePods) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sNodeAllocatablePods) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sNodeAllocatablePods) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sNodeAllocatablePods(cfg MetricConfig) metricK8sNodeAllocatablePods {
	m := metricK8sNodeAllocatablePods{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricK8sNodeAllocatableStorage struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.node.allocatable_storage metric with initial data.
func (m *metricK8sNodeAllocatableStorage) init() {
	m.data.SetName("k8s.node.allocatable_storage")
	m.data.SetDescription("How many bytes of storage remaining that the node can allocate to pods")
	m.data.SetUnit("By")
	m.data.SetEmptyGauge()
}

func (m *metricK8sNodeAllocatableStorage) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sNodeAllocatableStorage) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sNodeAllocatableStorage) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sNodeAllocatableStorage(cfg MetricConfig) metricK8sNodeAllocatableStorage {
	m := metricK8sNodeAllocatableStorage{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricK8sNodeConditionDiskPressure struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.node.condition_disk_pressure metric with initial data.
func (m *metricK8sNodeConditionDiskPressure) init() {
	m.data.SetName("k8s.node.condition_disk_pressure")
	m.data.SetDescription("Whether this node is DiskPressure (1), not DiskPressure (0) or in an unknown state (-1)")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricK8sNodeConditionDiskPressure) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sNodeConditionDiskPressure) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sNodeConditionDiskPressure) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sNodeConditionDiskPressure(cfg MetricConfig) metricK8sNodeConditionDiskPressure {
	m := metricK8sNodeConditionDiskPressure{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricK8sNodeConditionMemoryPressure struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.node.condition_memory_pressure metric with initial data.
func (m *metricK8sNodeConditionMemoryPressure) init() {
	m.data.SetName("k8s.node.condition_memory_pressure")
	m.data.SetDescription("Whether this node is MemoryPressure (1), not MemoryPressure (0) or in an unknown state (-1)")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricK8sNodeConditionMemoryPressure) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sNodeConditionMemoryPressure) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sNodeConditionMemoryPressure) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sNodeConditionMemoryPressure(cfg MetricConfig) metricK8sNodeConditionMemoryPressure {
	m := metricK8sNodeConditionMemoryPressure{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricK8sNodeConditionNetworkUnavailable struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.node.condition_network_unavailable metric with initial data.
func (m *metricK8sNodeConditionNetworkUnavailable) init() {
	m.data.SetName("k8s.node.condition_network_unavailable")
	m.data.SetDescription("Whether this node is NetworkUnavailable (1), not NetworkUnavailable (0) or in an unknown state (-1)")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricK8sNodeConditionNetworkUnavailable) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sNodeConditionNetworkUnavailable) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sNodeConditionNetworkUnavailable) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sNodeConditionNetworkUnavailable(cfg MetricConfig) metricK8sNodeConditionNetworkUnavailable {
	m := metricK8sNodeConditionNetworkUnavailable{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricK8sNodeConditionPidPressure struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.node.condition_pid_pressure metric with initial data.
func (m *metricK8sNodeConditionPidPressure) init() {
	m.data.SetName("k8s.node.condition_pid_pressure")
	m.data.SetDescription("Whether this node is PidPressure (1), not PidPressure (0) or in an unknown state (-1)")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricK8sNodeConditionPidPressure) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sNodeConditionPidPressure) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sNodeConditionPidPressure) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sNodeConditionPidPressure(cfg MetricConfig) metricK8sNodeConditionPidPressure {
	m := metricK8sNodeConditionPidPressure{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricK8sNodeConditionReady struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.node.condition_ready metric with initial data.
func (m *metricK8sNodeConditionReady) init() {
	m.data.SetName("k8s.node.condition_ready")
	m.data.SetDescription("Whether this node is Ready (1), not Ready (0) or in an unknown state (-1)")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricK8sNodeConditionReady) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sNodeConditionReady) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sNodeConditionReady) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sNodeConditionReady(cfg MetricConfig) metricK8sNodeConditionReady {
	m := metricK8sNodeConditionReady{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user config.
type MetricsBuilder struct {
	startTime                                pcommon.Timestamp   // start time that will be applied to all recorded data points.
	metricsCapacity                          int                 // maximum observed number of metrics per resource.
	metricsBuffer                            pmetric.Metrics     // accumulates metrics data before emitting.
	buildInfo                                component.BuildInfo // contains version information
	metricK8sNodeAllocatableCPU              metricK8sNodeAllocatableCPU
	metricK8sNodeAllocatableEphemeralStorage metricK8sNodeAllocatableEphemeralStorage
	metricK8sNodeAllocatableMemory           metricK8sNodeAllocatableMemory
	metricK8sNodeAllocatablePods             metricK8sNodeAllocatablePods
	metricK8sNodeAllocatableStorage          metricK8sNodeAllocatableStorage
	metricK8sNodeConditionDiskPressure       metricK8sNodeConditionDiskPressure
	metricK8sNodeConditionMemoryPressure     metricK8sNodeConditionMemoryPressure
	metricK8sNodeConditionNetworkUnavailable metricK8sNodeConditionNetworkUnavailable
	metricK8sNodeConditionPidPressure        metricK8sNodeConditionPidPressure
	metricK8sNodeConditionReady              metricK8sNodeConditionReady
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.CreateSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                                pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                            pmetric.NewMetrics(),
		buildInfo:                                settings.BuildInfo,
		metricK8sNodeAllocatableCPU:              newMetricK8sNodeAllocatableCPU(mbc.Metrics.K8sNodeAllocatableCPU),
		metricK8sNodeAllocatableEphemeralStorage: newMetricK8sNodeAllocatableEphemeralStorage(mbc.Metrics.K8sNodeAllocatableEphemeralStorage),
		metricK8sNodeAllocatableMemory:           newMetricK8sNodeAllocatableMemory(mbc.Metrics.K8sNodeAllocatableMemory),
		metricK8sNodeAllocatablePods:             newMetricK8sNodeAllocatablePods(mbc.Metrics.K8sNodeAllocatablePods),
		metricK8sNodeAllocatableStorage:          newMetricK8sNodeAllocatableStorage(mbc.Metrics.K8sNodeAllocatableStorage),
		metricK8sNodeConditionDiskPressure:       newMetricK8sNodeConditionDiskPressure(mbc.Metrics.K8sNodeConditionDiskPressure),
		metricK8sNodeConditionMemoryPressure:     newMetricK8sNodeConditionMemoryPressure(mbc.Metrics.K8sNodeConditionMemoryPressure),
		metricK8sNodeConditionNetworkUnavailable: newMetricK8sNodeConditionNetworkUnavailable(mbc.Metrics.K8sNodeConditionNetworkUnavailable),
		metricK8sNodeConditionPidPressure:        newMetricK8sNodeConditionPidPressure(mbc.Metrics.K8sNodeConditionPidPressure),
		metricK8sNodeConditionReady:              newMetricK8sNodeConditionReady(mbc.Metrics.K8sNodeConditionReady),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption func(pmetric.ResourceMetrics)

// WithResource sets the provided resource on the emitted ResourceMetrics.
// It's recommended to use ResourceBuilder to create the resource.
func WithResource(res pcommon.Resource) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		res.CopyTo(rm.Resource())
	}
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(rmo ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	rm.SetSchemaUrl(conventions.SchemaURL)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/k8sclusterreceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricK8sNodeAllocatableCPU.emit(ils.Metrics())
	mb.metricK8sNodeAllocatableEphemeralStorage.emit(ils.Metrics())
	mb.metricK8sNodeAllocatableMemory.emit(ils.Metrics())
	mb.metricK8sNodeAllocatablePods.emit(ils.Metrics())
	mb.metricK8sNodeAllocatableStorage.emit(ils.Metrics())
	mb.metricK8sNodeConditionDiskPressure.emit(ils.Metrics())
	mb.metricK8sNodeConditionMemoryPressure.emit(ils.Metrics())
	mb.metricK8sNodeConditionNetworkUnavailable.emit(ils.Metrics())
	mb.metricK8sNodeConditionPidPressure.emit(ils.Metrics())
	mb.metricK8sNodeConditionReady.emit(ils.Metrics())

	for _, op := range rmo {
		op(rm)
	}
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user config, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordK8sNodeAllocatableCPUDataPoint adds a data point to k8s.node.allocatable_cpu metric.
func (mb *MetricsBuilder) RecordK8sNodeAllocatableCPUDataPoint(ts pcommon.Timestamp, val float64) {
	mb.metricK8sNodeAllocatableCPU.recordDataPoint(mb.startTime, ts, val)
}

// RecordK8sNodeAllocatableEphemeralStorageDataPoint adds a data point to k8s.node.allocatable_ephemeral_storage metric.
func (mb *MetricsBuilder) RecordK8sNodeAllocatableEphemeralStorageDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricK8sNodeAllocatableEphemeralStorage.recordDataPoint(mb.startTime, ts, val)
}

// RecordK8sNodeAllocatableMemoryDataPoint adds a data point to k8s.node.allocatable_memory metric.
func (mb *MetricsBuilder) RecordK8sNodeAllocatableMemoryDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricK8sNodeAllocatableMemory.recordDataPoint(mb.startTime, ts, val)
}

// RecordK8sNodeAllocatablePodsDataPoint adds a data point to k8s.node.allocatable_pods metric.
func (mb *MetricsBuilder) RecordK8sNodeAllocatablePodsDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricK8sNodeAllocatablePods.recordDataPoint(mb.startTime, ts, val)
}

// RecordK8sNodeAllocatableStorageDataPoint adds a data point to k8s.node.allocatable_storage metric.
func (mb *MetricsBuilder) RecordK8sNodeAllocatableStorageDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricK8sNodeAllocatableStorage.recordDataPoint(mb.startTime, ts, val)
}

// RecordK8sNodeConditionDiskPressureDataPoint adds a data point to k8s.node.condition_disk_pressure metric.
func (mb *MetricsBuilder) RecordK8sNodeConditionDiskPressureDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricK8sNodeConditionDiskPressure.recordDataPoint(mb.startTime, ts, val)
}

// RecordK8sNodeConditionMemoryPressureDataPoint adds a data point to k8s.node.condition_memory_pressure metric.
func (mb *MetricsBuilder) RecordK8sNodeConditionMemoryPressureDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricK8sNodeConditionMemoryPressure.recordDataPoint(mb.startTime, ts, val)
}

// RecordK8sNodeConditionNetworkUnavailableDataPoint adds a data point to k8s.node.condition_network_unavailable metric.
func (mb *MetricsBuilder) RecordK8sNodeConditionNetworkUnavailableDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricK8sNodeConditionNetworkUnavailable.recordDataPoint(mb.startTime, ts, val)
}

// RecordK8sNodeConditionPidPressureDataPoint adds a data point to k8s.node.condition_pid_pressure metric.
func (mb *MetricsBuilder) RecordK8sNodeConditionPidPressureDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricK8sNodeConditionPidPressure.recordDataPoint(mb.startTime, ts, val)
}

// RecordK8sNodeConditionReadyDataPoint adds a data point to k8s.node.condition_ready metric.
func (mb *MetricsBuilder) RecordK8sNodeConditionReadyDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricK8sNodeConditionReady.recordDataPoint(mb.startTime, ts, val)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
