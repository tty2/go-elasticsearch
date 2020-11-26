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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search.asciidoc#L96>
//
// --------------------------------------------------------------------------------
// POST /_search
// {
//   "query" : {
//     "match_all" : {}
//   },
//   "stats" : ["group1", "group2"]
// }
// --------------------------------------------------------------------------------

func Test_search_1284dcc6c5bee851d2e882fb6907537b(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:1284dcc6c5bee851d2e882fb6907537b[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "match_all": {}
		  },
		  "stats": [
		    "group1",
		    "group2"
		  ]
		}`)),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:1284dcc6c5bee851d2e882fb6907537b[]
}
