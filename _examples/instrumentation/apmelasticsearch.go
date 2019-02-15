package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/elastic/go-elasticsearch"

	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmelasticsearch"
)

const count = 100

var (
	tWidth, _, _ = terminal.GetSize(int(os.Stdout.Fd()))

	faint   = color.New(color.Faint)
	bold    = color.New(color.Bold)
	boldRed = color.New(color.FgRed).Add(color.Bold)
)

func init() {
	if tWidth < 0 {
		tWidth = 0
	}
}

func main() {
	log.SetFlags(0)
	start := time.Now()

	// Create new elasticsearch client ...
	//
	es, err := elasticsearch.NewClient(
		elasticsearch.Config{
			// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
			// ... using the "apmelasticsearch" wrapper for instrumentation
			Transport: apmelasticsearch.WrapRoundTripper(http.DefaultTransport),
			// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
		})
	if err != nil {
		log.Fatal("ERROR: %s", err)
	}

	// Set up the "done" channel
	//
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt)

	// Set up tickers
	//
	tickers := struct {
		Info   *time.Ticker
		Index  *time.Ticker
		Health *time.Ticker
		Search *time.Ticker
	}{
		Info:   time.NewTicker(time.Second),
		Index:  time.NewTicker(500 * time.Millisecond),
		Health: time.NewTicker(5 * time.Second),
		Search: time.NewTicker(10 * time.Second),
	}
	defer tickers.Info.Stop()
	defer tickers.Index.Stop()
	defer tickers.Health.Stop()
	defer tickers.Search.Stop()

	// Initialize the context
	//
	ctx := context.Background()

	// Perform API calls
	//
	for {
		select {
		case <-done:
			fmt.Print("\n")
			fmt.Println(strings.Repeat("━", tWidth))
			faint.Printf("Finished in %s\n\n", time.Now().Sub(start).Truncate(time.Second))
			return

		// -> Info
		//
		case <-tickers.Info.C:
			// Create a transaction and span
			// TODO(karmi): Shouldn't be needed?
			tp := apm.DefaultTracer.StartTransaction("GET /", "request")
			sp, ctx := apm.StartSpan(ctx, "Info", "elasticsearch")

			res, err := es.Info(es.Info.WithContext(ctx))
			if err != nil {
				boldRed.Printf("Error getting response: %s\n", err)
				apm.DefaultTracer.NewError(err).Send()
			} else {
				faint.Println(res.Status())
				res.Body.Close()
			}
			tp.End()
			sp.End()

		// -> Index
		//
		case t := <-tickers.Index.C:
			// Artificially fail some requests...
			var body io.Reader
			if t.Second()%4 == 0 {
				body = strings.NewReader(``)
			} else {
				body = strings.NewReader(`{"timestamp":"` + t.Format(time.RFC3339) + `"}`)
			}

			res, err := es.Index("test", body, es.Index.WithContext(ctx))
			if err != nil {
				boldRed.Printf("Error getting response: %s\n", err)
			} else {
				faint.Println(res.Status())
				res.Body.Close()
			}

		// -> Health
		//
		case <-tickers.Health.C:
			res, err := es.Cluster.Health(
				es.Cluster.Health.WithLevel("indices"),
				es.Cluster.Health.WithContext(ctx),
			)
			if err != nil {
				boldRed.Printf("Error getting response: %s\n", err)
			} else {
				faint.Println(res.Status())
				res.Body.Close()
			}

		// -> Search
		//
		case <-tickers.Search.C:
			res, err := es.Search(
				es.Search.WithIndex("test"),
				es.Search.WithSort("timestamp:desc"),
				es.Search.WithSize(1),
				es.Search.WithContext(ctx),
			)
			if err != nil {
				boldRed.Printf("Error getting response: %s\n", err)
			} else {
				faint.Println(res.Status())
				res.Body.Close()
			}
		}
	}
}
