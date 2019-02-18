package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
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
			// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
			// Set up the APM transaction and put it into the context
			tpx := apm.DefaultTracer.StartTransaction("Info()", "monitoring")
			ctx := apm.ContextWithTransaction(ctx, tpx)
			// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

			res, err := es.Info(es.Info.WithContext(ctx))
			if err != nil {
				boldRed.Printf("Error getting response: %s\n", err)
				apm.CaptureError(ctx, err).Send()
				break
			}

			faint.Println(res.Status())
			res.Body.Close()

			// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
			// Send the transaction to the APM server
			tpx.End()
			// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

		// -> Index
		//
		case t := <-tickers.Index.C:
			// Fail some requests with empty body...
			var body io.Reader
			if t.Second()%4 == 0 {
				body = strings.NewReader(``)
			} else {
				body = strings.NewReader(`{"timestamp":"` + t.Format(time.RFC3339) + `"}`)
			}

			tpx := apm.DefaultTracer.StartTransaction("Index()", "indexing")
			ctx := apm.ContextWithTransaction(ctx, tpx)

			res, err := es.Index("test", body, es.Index.WithContext(ctx))
			if err != nil {
				boldRed.Printf("Error getting response: %s\n", err)
				apm.CaptureError(ctx, err).Send()
				break
			}

			faint.Println(res.Status())
			res.Body.Close()
			tpx.End()

		// -> Health
		//
		case <-tickers.Health.C:
			tpx := apm.DefaultTracer.StartTransaction("Cluster.Health()", "monitoring")
			ctx := apm.ContextWithTransaction(ctx, tpx)

			res, err := es.Cluster.Health(
				es.Cluster.Health.WithLevel("indices"),
				es.Cluster.Health.WithContext(ctx),
			)
			if err != nil {
				boldRed.Printf("Error getting response: %s\n", err)
				apm.CaptureError(ctx, err).Send()
				break
			}

			faint.Println(res.Status())
			res.Body.Close()
			tpx.End()

		// -> Search
		//
		case <-tickers.Search.C:
			tpx := apm.DefaultTracer.StartTransaction("Search()", "searching")
			ctx := apm.ContextWithTransaction(ctx, tpx)

			// Randomly trigger an error
			if rand.Intn(10) > 8 {
				err := errors.New("Custom error")
				boldRed.Println("ERROR:", err)
				apm.CaptureError(ctx, err).Send()
				break
			}

			res, err := es.Search(
				es.Search.WithIndex("test"),
				es.Search.WithSort("timestamp:desc"),
				es.Search.WithSize(1),
				es.Search.WithContext(ctx),
			)
			if err != nil {
				boldRed.Printf("Error getting response: %s\n", err)
				apm.CaptureError(ctx, err).Send()
				break
			}

			// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
			// Create a custom span within the transaction
			sp := tpx.StartSpan("JSON/Decode", "searching", nil)
			// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

			var r map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
				sp.End()
				boldRed.Printf("Error parsing the response body: %s\n", err)
				apm.CaptureError(ctx, err).Send()
				break
			}
			sp.End()

			// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
			// Create a custom span within the transaction
			sp = tpx.StartSpan("UI/Render", "searching", nil)
			// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
			faint.Printf("%s ; %vms\n", res.Status(), r["took"])
			sp.End()

			res.Body.Close()
			tpx.End()
		}
	}
}
