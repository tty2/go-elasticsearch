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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/mapping/types/nested.asciidoc#L22>
//
// --------------------------------------------------------------------------------
// PUT my-index-000001/_doc/1
// {
//   "group" : "fans",
//   "user" : [ <1>
//     {
//       "first" : "John",
//       "last" :  "Smith"
//     },
//     {
//       "first" : "Alice",
//       "last" :  "White"
//     }
//   ]
// }
// --------------------------------------------------------------------------------

func Test_mapping_types_nested_7fef68840761c6982c14ad7af96caf37(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:7fef68840761c6982c14ad7af96caf37[]
	res, err := es.Index(
		"my-index-000001",
		strings.NewReader(`{
		  "group": "fans",
		  "user": [
		    {
		      "first": "John",
		      "last": "Smith"
		    },
		    {
		      "first": "Alice",
		      "last": "White"
		    }
		  ]
		}`),
		es.Index.WithDocumentID("1"),
		es.Index.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:7fef68840761c6982c14ad7af96caf37[]
}
