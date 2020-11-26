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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/match-query.asciidoc#L219>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//   "query": {
//     "match": {
//       "message": {
//         "query": "this is a testt",
//         "fuzziness": "AUTO"
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_match_query_6b0288acb739c4667d41339e5100c327(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:6b0288acb739c4667d41339e5100c327[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "match": {
		      "message": {
		        "query": "this is a testt",
		        "fuzziness": "AUTO"
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
	// end:6b0288acb739c4667d41339e5100c327[]
}
