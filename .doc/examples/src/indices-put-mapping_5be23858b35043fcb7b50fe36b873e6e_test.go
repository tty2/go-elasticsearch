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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/put-mapping.asciidoc#L13>
//
// --------------------------------------------------------------------------------
// PUT /twitter/_mapping
// {
//   "properties": {
//     "email": {
//       "type": "keyword"
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_indices_put_mapping_5be23858b35043fcb7b50fe36b873e6e(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:5be23858b35043fcb7b50fe36b873e6e[]
	res, err := es.Indices.PutMapping(
		[]string{"twitter"},
		strings.NewReader(`{
		  "properties": {
		    "email": {
		      "type": "keyword"
		    }
		  }
		}`),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:5be23858b35043fcb7b50fe36b873e6e[]
}
