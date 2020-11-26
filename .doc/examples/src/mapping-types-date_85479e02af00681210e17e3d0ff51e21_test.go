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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/mapping/types/date.asciidoc#L77>
//
// --------------------------------------------------------------------------------
// PUT my-index-000001
// {
//   "mappings": {
//     "properties": {
//       "date": {
//         "type":   "date",
//         "format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_mapping_types_date_85479e02af00681210e17e3d0ff51e21(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:85479e02af00681210e17e3d0ff51e21[]
	res, err := es.Indices.Create(
		"my-index-000001",
		es.Indices.Create.WithBody(strings.NewReader(`{
		  "mappings": {
		    "properties": {
		      "date": {
		        "type": "date",
		        "format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
		      }
		    }
		  }
		}`)),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:85479e02af00681210e17e3d0ff51e21[]
}
