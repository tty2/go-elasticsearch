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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/reindex.asciidoc#L850>
//
// --------------------------------------------------------------------------------
// POST _reindex
// {
//   "source": {
//     "index": "twitter"
//   },
//   "dest": {
//     "index": "new_twitter",
//     "version_type": "external"
//   },
//   "script": {
//     "source": "if (ctx._source.foo == 'bar') {ctx._version++; ctx._source.remove('foo')}",
//     "lang": "painless"
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_reindex_8871b8fcb6de4f0c7dff22798fb10fb7(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:8871b8fcb6de4f0c7dff22798fb10fb7[]
	res, err := es.Reindex(
		strings.NewReader(`{
		  "source": {
		    "index": "twitter"
		  },
		  "dest": {
		    "index": "new_twitter",
		    "version_type": "external"
		  },
		  "script": {
		    "source": "if (ctx._source.foo == 'bar') {ctx._version++; ctx._source.remove('foo')}",
		    "lang": "painless"
		  }
		}`))
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:8871b8fcb6de4f0c7dff22798fb10fb7[]
}
