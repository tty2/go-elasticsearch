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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/api-conventions.asciidoc#L282>
//
// --------------------------------------------------------------------------------
// GET /_cluster/state?filter_path=routing_table.indices.**.state
// --------------------------------------------------------------------------------

func Test_api_conventions_1252fa45847edba5ec2b2f33da70ec5b(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:1252fa45847edba5ec2b2f33da70ec5b[]
	res, err := es.Cluster.State(
		es.Cluster.State.WithFilterPath("routing_table.indices.**.state"),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:1252fa45847edba5ec2b2f33da70ec5b[]
}
