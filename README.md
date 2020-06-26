# hcl-opentelemetry
This library provides an `otel` package to load HCL configuration and initialize OpenTelemetry sinks to use for metrics reporting.

## Sinks
* Stdout
* Dogstatsd
* Prometheus

## Configuration
The library supports configuration files written in the [HashiCorp Configuration Language](https://github.com/hashicorp/hcl), and by proxy this means it is also JSON compatible.

```hcl
telemetry {
  stdout {
    period = "60s"
    pretty_print = false
    do_not_print_time = false
  }
}
```

```json
{
  "telemetry": {
    "stdout": {
      "period": "60s",
      "pretty_print": false,
      "do_not_print_time": false
    }
  }
}
```
