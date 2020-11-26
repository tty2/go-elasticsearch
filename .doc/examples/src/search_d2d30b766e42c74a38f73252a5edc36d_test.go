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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search.asciidoc#L32>
//
// --------------------------------------------------------------------------------
// POST /twitter/_search?routing=kimchy
// {
//   "query": {
//     "bool": {
//       "must": {
//         "query_string": {
//           "query": "some query string here"
//         }
//       },
//       "filter": {
//         "term": { "user": "kimchy" }
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_search_d2d30b766e42c74a38f73252a5edc36d(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:d2d30b766e42c74a38f73252a5edc36d[]
	res, err := es.Search(
		es.Search.WithIndex("twitter"),
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "bool": {
		      "must": {
		        "query_string": {
		          "query": "some query string here"
		        }
		      },
		      "filter": {
		        "term": {
		          "user": "kimchy"
		        }
		      }
		    }
		  }
		}`)),
		es.Search.WithRouting("kimchy"),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:d2d30b766e42c74a38f73252a5edc36d[]
}
