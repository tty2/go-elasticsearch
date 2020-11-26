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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/mapping/types/date.asciidoc#L35>
//
// --------------------------------------------------------------------------------
// PUT my-index-000001
// {
//   "mappings": {
//     "properties": {
//       "date": {
//         "type": "date" <1>
//       }
//     }
//   }
// }
//
// PUT my-index-000001/_doc/1
// { "date": "2015-01-01" } <2>
//
// PUT my-index-000001/_doc/2
// { "date": "2015-01-01T12:10:30Z" } <3>
//
// PUT my-index-000001/_doc/3
// { "date": 1420070400001 } <4>
//
// GET my-index-000001/_search
// {
//   "sort": { "date": "asc"} <5>
// }
// --------------------------------------------------------------------------------

func Test_mapping_types_date_d2f52c106685bd8eab47e11d644d7a70(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:d2f52c106685bd8eab47e11d644d7a70[]
	{
		res, err := es.Indices.Create(
			"my-index-000001",
			es.Indices.Create.WithBody(strings.NewReader(`{
		  "mappings": {
		    "properties": {
		      "date": {
		        "type": "date"
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
	}

	{
		res, err := es.Index(
			"my-index-000001",
			strings.NewReader(`{
		  "date": "2015-01-01"
		} `),
			es.Index.WithDocumentID("1"),
			es.Index.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Index(
			"my-index-000001",
			strings.NewReader(`{
		  "date": "2015-01-01T12:10:30Z"
		} `),
			es.Index.WithDocumentID("2"),
			es.Index.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Index(
			"my-index-000001",
			strings.NewReader(`{
		  "date": 1420070400001
		} `),
			es.Index.WithDocumentID("3"),
			es.Index.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Search(
			es.Search.WithIndex("my-index-000001"),
			es.Search.WithBody(strings.NewReader(`{
		  "sort": {
		    "date": "asc"
		  }
		}`)),
			es.Search.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}
	// end:d2f52c106685bd8eab47e11d644d7a70[]
}
