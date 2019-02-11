package main

import (
	"context"
	"log"
	"strings"
	"time"

	"go.opencensus.io/examples/exporter"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"go.opencensus.io/trace"

	"github.com/elastic/go-elasticsearch"
)

const count = 100

var (
	tMethod, _ = tag.NewKey("method")

	// TEMP > Prevent compilation errors
	_, _, _ = stats.Record, trace.RegisterExporter, tag.New
)

func main() {
	log.SetFlags(0)

	// Register views bundled with the "ochttp" plugin
	//
	if err := view.Register(
		ochttp.ClientRoundtripLatencyDistribution,
		ochttp.ClientCompletedCount,
	); err != nil {
		log.Fatal("ERROR: %s", err)
	}

	// Report views to STDOUT every second
	//
	view.SetReportingPeriod(time.Second)
	view.RegisterExporter(&exporter.PrintExporter{})

	// Report 5% of traces to STDOUT
	//
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.ProbabilitySampler(0.05)})
	trace.RegisterExporter(&exporter.PrintExporter{})

	// Create new elasticsearch client ...
	//
	es, err := elasticsearch.NewClient(
		elasticsearch.Config{
			// ... using the "ochttp" wrapper
			Transport: &ochttp.Transport{},
		})
	if err != nil {
		log.Fatal("ERROR: %s", err)
	}

	ctx, _ := tag.New(context.Background(), tag.Upsert(tMethod, "main"))
	for i := 0; i < count; i++ {
		// Create span
		ctx, span := trace.StartSpan(ctx, "es/Info")

		// Call Elasticsearch API
		res, err := es.Info(es.Info.WithContext(ctx))
		if err != nil {
			log.Printf("Error getting response: %s", err)
		}
		// NOTE: Need to close the body so the stats are recorded by ochttp
		res.Body.Close()
		span.End()
	}

	// Wait a bit for the collection and reporting to finish.
	//
	time.Sleep(1250 * time.Millisecond)
	log.Println(strings.Repeat("=", 80))
}
