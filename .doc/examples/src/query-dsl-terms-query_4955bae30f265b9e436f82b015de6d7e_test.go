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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/terms-query.asciidoc#L186>
//
// --------------------------------------------------------------------------------
// GET my-index-000001/_search?pretty
// {
//   "query": {
//     "terms": {
//         "color" : {
//             "index" : "my-index-000001",
//             "id" : "2",
//             "path" : "color"
//         }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_terms_query_4955bae30f265b9e436f82b015de6d7e(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:4955bae30f265b9e436f82b015de6d7e[]
	res, err := es.Search(
		es.Search.WithIndex("my-index-000001"),
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "terms": {
		      "color": {
		        "index": "my-index-000001",
		        "id": "2",
		        "path": "color"
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
	// end:4955bae30f265b9e436f82b015de6d7e[]
}
