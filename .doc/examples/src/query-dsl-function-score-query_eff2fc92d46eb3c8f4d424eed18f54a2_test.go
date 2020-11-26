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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/function-score-query.asciidoc#L19>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//   "query": {
//     "function_score": {
//       "query": { "match_all": {} },
//       "boost": "5",
//       "random_score": {}, <1>
//       "boost_mode": "multiply"
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_function_score_query_eff2fc92d46eb3c8f4d424eed18f54a2(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:eff2fc92d46eb3c8f4d424eed18f54a2[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "function_score": {
		      "query": {
		        "match_all": {}
		      },
		      "boost": "5",
		      "random_score": {},
		      "boost_mode": "multiply"
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
	// end:eff2fc92d46eb3c8f4d424eed18f54a2[]
}
