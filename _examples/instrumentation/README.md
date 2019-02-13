# Example: Instrumentation

This example demonstrates how to instrument the Elasticsearch client.

### OpenCensus

The [**`opencensus`**](./opencensus) example uses the [`ochttp.Transport`](https://godoc.org/go.opencensus.io/plugin/ochttp#example-Transport) wrapper to auto-instrument the client calls, and provides a simple exporter which prints information to the terminal.

[![asciicast](https://asciinema.org/a/KhyP3GuuHPJAZQAmrgmdwS7uf.svg)](https://asciinema.org/a/KhyP3GuuHPJAZQAmrgmdwS7uf)
