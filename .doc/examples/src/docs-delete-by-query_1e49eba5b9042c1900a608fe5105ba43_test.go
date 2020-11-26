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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/delete-by-query.asciidoc#L421>
//
// --------------------------------------------------------------------------------
// POST twitter/_delete_by_query
// {
//   "slice": {
//     "id": 0,
//     "max": 2
//   },
//   "query": {
//     "range": {
//       "likes": {
//         "lt": 10
//       }
//     }
//   }
// }
// POST twitter/_delete_by_query
// {
//   "slice": {
//     "id": 1,
//     "max": 2
//   },
//   "query": {
//     "range": {
//       "likes": {
//         "lt": 10
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_delete_by_query_1e49eba5b9042c1900a608fe5105ba43(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:1e49eba5b9042c1900a608fe5105ba43[]
	{
		res, err := es.DeleteByQuery(
			[]string{"twitter"},
			strings.NewReader(`{
		  "slice": {
		    "id": 0,
		    "max": 2
		  },
		  "query": {
		    "range": {
		      "likes": {
		        "lt": 10
		      }
		    }
		  }
		}`),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.DeleteByQuery(
			[]string{"twitter"},
			strings.NewReader(`{
		  "slice": {
		    "id": 1,
		    "max": 2
		  },
		  "query": {
		    "range": {
		      "likes": {
		        "lt": 10
		      }
		    }
		  }
		}`),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}
	// end:1e49eba5b9042c1900a608fe5105ba43[]
}
