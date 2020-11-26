module github.com/elastic/go-elasticsearch/v8/_examples/bulk/kafka

go 1.11

replace github.com/elastic/go-elasticsearch/v8 => ../../..

require (
	github.com/dustin/go-humanize v1.0.0
	github.com/elastic/go-elasticsearch/v8 master
	github.com/segmentio/kafka-go v0.3.5
	go.elastic.co/apm v1.7.1
	go.elastic.co/apm/module/apmelasticsearch v1.7.1
)
