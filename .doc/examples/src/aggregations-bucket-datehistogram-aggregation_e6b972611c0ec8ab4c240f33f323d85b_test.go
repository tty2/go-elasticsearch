// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated, DO NOT EDIT

package elasticsearch_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/aggregations/bucket/datehistogram-aggregation.asciidoc#L431>
//
// --------------------------------------------------------------------------------
// GET my-index-000001/_search?size=0
// {
//   "aggs": {
//     "by_day": {
//       "date_histogram": {
//         "field":     "date",
//         "calendar_interval":  "day",
//         "time_zone": "-01:00"
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_aggregations_bucket_datehistogram_aggregation_e6b972611c0ec8ab4c240f33f323d85b(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:e6b972611c0ec8ab4c240f33f323d85b[]
	res, err := es.Search(
		es.Search.WithIndex("my-index-000001"),
		es.Search.WithBody(strings.NewReader(`{
		  "aggs": {
		    "by_day": {
		      "date_histogram": {
		        "field": "date",
		        "calendar_interval": "day",
		        "time_zone": "-01:00"
		      }
		    }
		  }
		}`)),
		es.Search.WithSize(0),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:e6b972611c0ec8ab4c240f33f323d85b[]
}
