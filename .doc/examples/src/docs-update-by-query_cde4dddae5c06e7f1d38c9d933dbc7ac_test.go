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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/update-by-query.asciidoc#L326>
//
// --------------------------------------------------------------------------------
// POST twitter,blog/_update_by_query
// --------------------------------------------------------------------------------

func Test_docs_update_by_query_cde4dddae5c06e7f1d38c9d933dbc7ac(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:cde4dddae5c06e7f1d38c9d933dbc7ac[]
	res, err := es.UpdateByQuery(
		[]string{"twitter,blog"},
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:cde4dddae5c06e7f1d38c9d933dbc7ac[]
}
