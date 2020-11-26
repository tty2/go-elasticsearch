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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/reindex.asciidoc#L716>
//
// --------------------------------------------------------------------------------
// POST _reindex
// {
//   "source": {
//     "index": "test"
//   },
//   "dest": {
//     "index": "test2"
//   },
//   "script": {
//     "source": "ctx._source.tag = ctx._source.remove(\"flag\")"
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_reindex_1216f8f7367df3aa823012cef310c08a(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:1216f8f7367df3aa823012cef310c08a[]
	res, err := es.Reindex(
		strings.NewReader(`{
		  "source": {
		    "index": "test"
		  },
		  "dest": {
		    "index": "test2"
		  },
		  "script": {
		    "source": "ctx._source.tag = ctx._source.remove(\"flag\")"
		  }
		}`))
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:1216f8f7367df3aa823012cef310c08a[]
}
