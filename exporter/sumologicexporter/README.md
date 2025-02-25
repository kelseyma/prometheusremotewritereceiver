# Sumo Logic Exporter

<!-- status autogenerated section -->
| Status        |           |
| ------------- |-----------|
| Stability     | [beta]: metrics, logs   |
| Distributions | [contrib] |
| Issues        | ![Open issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aopen%20label%3Aexporter%2Fsumologic%20&label=open&color=orange&logo=opentelemetry) ![Closed issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aclosed%20label%3Aexporter%2Fsumologic%20&label=closed&color=blue&logo=opentelemetry) |
| [Code Owners](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/CONTRIBUTING.md#becoming-a-code-owner)    | [@sumo-drosiek](https://www.github.com/sumo-drosiek) |

[beta]: https://github.com/open-telemetry/opentelemetry-collector#beta
[contrib]: https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol-contrib
<!-- end autogenerated section -->

## Migration to new architecture

**This exporter is undergoing major changes right now.**

For some time we have been developing the [new Sumo Logic exporter](https://github.com/SumoLogic/sumologic-otel-collector/tree/main/pkg/exporter/sumologicexporter#sumo-logic-exporter) and now we are in the process of moving it into this repository.

The following options are deprecated and they will not exist in the new version:

- `metric_format: {carbon2, graphite}`
- `metadata_attributes: [<regex>]`
- `graphite_template: <template>`
- `source_category: <template>`
- `source_name: <template>`
- `source_host: <template>`

After the new exporter will be moved to this repository:

- `carbon2` and `graphite` are going to be no longer supported and `prometheus` or `otlp` format should be used
- all resource level attributes are going to be treated as `metadata_attributes`. You can use [Group by Attributes processor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/groupbyattrsprocessor) to move attributes from record level to resource level. For example:

  ```yaml
  # before switch to new collector
  exporters:
    sumologic:
      metadata_attribute:
        - my_attribute
  # after switch to new collector
  processors:
    groupbyattrs:
      keys:
        - my_attribute
  ```

- Source templates (`source_category`, `source_name` and `source_host`) are going to be removed from the exporter and sources may be set using `_sourceCategory`, `sourceName` or `_sourceHost` resource attributes. We recommend to use [Transform Processor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/transformprocessor/). For example:

  ```yaml
  # before switch to new collector
  exporters:
    sumologic:
      source_category: "%{foo}/constant/%{bar}"
  # after switch to new collector
  processors:
    transformprocessor:
      log_statements:
        context: log
        statements:
          # set default value to unknown
          - set(attributes["foo"], "unknown") where attributes["foo"] == nil
          - set(attributes["bar"], "unknown") where attributes["foo"] == nil
          # set _sourceCategory as "%{foo}/constant/%{bar}"
          - set(resource.attributes["_sourceCategory"], Concat([attributes["foo"], "/constant/", attributes["bar"]], ""))
  ```

## Configuration

This exporter supports sending logs and metrics data to [Sumo Logic](https://www.sumologic.com/).
Traces are exported using native otlphttp exporter as described
[here](https://help.sumologic.com/Traces/Getting_Started_with_Transaction_Tracing)

Configuration is specified via the yaml in the following structure:

```yaml
exporters:
  # ...
  sumologic:
    # unique URL generated for your HTTP Source, this is the address to send data to
    endpoint: <HTTP_Source_URL>
    # Compression encoding format, empty string means no compression, default = gzip
    compress_encoding: {gzip, deflate, ""}
    # max HTTP request body size in bytes before compression (if applied),
    # default = 1_048_576 (1MB)
    max_request_body_size: <max_request_body_size>

    # List of regexes for attributes which should be send as metadata
    # default = []
    #
    # This option is deprecated:
    # https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/sumologicexporter#migration-to-new-architecture
    metadata_attributes: [<regex>]

    # format to use when sending logs to Sumo Logic, default = json,
    log_format: {json, text}

    # format to use when sending metrics to Sumo Logic, default = prometheus,
    #
    # carbon2 and graphite are deprecated:
    # https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/sumologicexporter#migration-to-new-architecture
    metric_format: {carbon2, graphite, prometheus}

    # Template for Graphite format.
    # this option affects graphite format only
    # By default this is "%{_metric_}".
    #
    # Please regfer to Source temmplates for formatting explanation:
    # https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/sumologicexporter#source-templates
    #
    # This option is deprecated:
    # https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/sumologicexporter#migration-to-new-architecture
    graphite_template: <template>

    # Desired source category. Useful if you want to override the source category configured for the source.
    #
    # Please regfer to Source temmplates for formatting explanation:
    # https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/sumologicexporter#source-templates
    #
    # This option is deprecated:
    # https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/sumologicexporter#migration-to-new-architecture
    source_category: <template>

    # Desired source name. Useful if you want to override the source name configured for the source.
    #
    # Please regfer to Source temmplates for formatting explanation:
    # https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/sumologicexporter#source-templates
    #
    # This option is deprecated:
    # https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/sumologicexporter#migration-to-new-architecture
    source_name: <template>

    # Desired source host. Useful if you want to override the source hosy configured for the source.
    #
    # Please regfer to Source temmplates for formatting explanation:
    # https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/sumologicexporter#source-templates
    #
    # This option is deprecated:
    # https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/sumologicexporter#migration-to-new-architecture
    source_host: <template>

    # timeout is the timeout for every attempt to send data to the backend,
    # maximum connection timeout is 55s, default = 5s
    timeout: <timeout>

    # for below described queueing and retry related configuration please refer to:
    # https://github.com/open-telemetry/opentelemetry-collector/blob/main/exporter/exporterhelper/README.md#configuration

    retry_on_failure:
      # default = true
      enabled: {true, false}
      # time to wait after the first failure before retrying;
      # ignored if enabled is false, default = 5s
      initial_interval: <initial_interval>
      # is the upper bound on backoff; ignored if enabled is false, default = 30s
      max_interval: <max_interval>
      # is the maximum amount of time spent trying to send a batch;
      # ignored if enabled is false, default = 120s
      max_elapsed_time: <max_elapsed_time>

    sending_queue:
      # default = false
      enabled: {true, false}
      # number of consumers that dequeue batches; ignored if enabled is false,
      # default = 10
      num_consumers: <num_consumers>
      # when set, enables persistence and uses the component specified as a storage extension for the persistent queue
      # make sure to configure and add a `file_storage` extension in `service.extensions`.
      # default = None
      storage: <storage_name>
      # maximum number of batches kept in memory before data;
      # ignored if enabled is false, default = 1000
      #
      # user should calculate this as num_seconds * requests_per_second where:
      # num_seconds is the number of seconds to buffer in case of a backend outage,
      # requests_per_second is the average number of requests per seconds.
      queue_size: <queue_size>
```

## Source Templates

You can specify a template with an attribute for `source_category`, `source_name`, `source_host` or `graphite_template` using `%{attr_name}`.

For example, when there is an attribute `my_attr`: `my_value`, `metrics/%{my_attr}` would be expanded to `metrics/my_value`.

For `graphite_template`, in addition to above, `%{_metric_}` is going to be replaced with metric name.

## Example Configuration

```yaml
exporters:
  sumologic:
    endpoint: http://localhost:3000
    compress_encoding: "gzip"
    max_request_body_size: "1_048_576"  # 1MB
    log_format: "text"
    metric_format: "prometheus"
    source_category: "custom category"
    source_name: "custom name"
    source_host: "custom host"
    metadata_attributes:
      - k8s.*
```
