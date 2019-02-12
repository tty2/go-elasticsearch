module github.com/elastic/go-elasticsearch/_examples/instrumentation/opencensus

replace github.com/elastic/go-elasticsearch => ../..

require (
	github.com/elastic/go-elasticsearch v0.0.0

	github.com/fatih/color v1.7.0
	github.com/mattn/go-colorable v0.1.0 // indirect
	github.com/mattn/go-isatty v0.0.4 // indirect
	go.opencensus.io v0.19.0
	golang.org/x/crypto v0.0.0-20190211182817-74369b46fc67
)
