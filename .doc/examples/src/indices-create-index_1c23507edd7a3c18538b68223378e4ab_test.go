// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated, DO NOT EDIT

package elasticsearch_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/create-index.asciidoc#L10>
//
// --------------------------------------------------------------------------------
// PUT /twitter
// --------------------------------------------------------------------------------

func Test_indices_create_index_1c23507edd7a3c18538b68223378e4ab(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:1c23507edd7a3c18538b68223378e4ab[]
	res, err := es.Indices.Create("twitter")
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:1c23507edd7a3c18538b68223378e4ab[]
}
