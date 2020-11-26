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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/range-query.asciidoc#L16>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//   "query": {
//     "range": {
//       "age": {
//         "gte": 10,
//         "lte": 20,
//         "boost": 2.0
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_range_query_6b77795e9249c8d9865f7a49fd86a863(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:6b77795e9249c8d9865f7a49fd86a863[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "range": {
		      "age": {
		        "gte": 10,
		        "lte": 20,
		        "boost": 2.0
		      }
		    }
		  }
		}`)),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:6b77795e9249c8d9865f7a49fd86a863[]
}
