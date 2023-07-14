# Prometheus Remote Write Receiver

| Status                   |                 |
|--------------------------|-----------------|
| Stability                | [development] |
| Supported pipeline types | metrics         |
| Distributions            | [contrib]       |

Supported pipeline types: metrics

## Getting Started

All that is required to enable the Prometheus Remote Write receiver is to include it in the
receiver definitions.

```yaml
receivers:
  prometheusremotewrite:
```

Http Server configuration settings:

- `endpoint` (default = 0.0.0.0:19291): host:port to which the receiver is going
  to receive data.

Additional server settings are mentioned here:
<https://github.com/open-telemetry/opentelemetry-collector/tree/main/config/confighttp#server-configuration>

- `time_threshold` (default = 24) - time threshold (in hours). All `timeseries` older than limit will be dropped.



[development]: https://github.com/open-telemetry/opentelemetry-collector#development
[contrib]: https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol-contrib