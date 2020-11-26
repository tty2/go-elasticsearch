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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/update-settings.asciidoc#L169>
//
// --------------------------------------------------------------------------------
// POST /twitter/_close
//
// PUT /twitter/_settings
// {
//   "analysis" : {
//     "analyzer":{
//       "content":{
//         "type":"custom",
//         "tokenizer":"whitespace"
//       }
//     }
//   }
// }
//
// POST /twitter/_open
// --------------------------------------------------------------------------------

func Test_indices_update_settings_ba0b4081c98f3387f76b77847c52ee9a(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:ba0b4081c98f3387f76b77847c52ee9a[]
	{
		res, err := es.Indices.Close([]string{"twitter"})
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Indices.PutSettings(strings.NewReader(`{
		  "analysis": {
		    "analyzer": {
		      "content": {
		        "type": "custom",
		        "tokenizer": "whitespace"
		      }
		    }
		  }
		}`),

			es.Indices.PutSettings.WithIndex([]string{"twitter"}...),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Indices.Open([]string{"twitter"})
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}
	// end:ba0b4081c98f3387f76b77847c52ee9a[]
}
