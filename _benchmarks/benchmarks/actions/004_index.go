// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package actions

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/tidwall/gjson"

	"github.com/elastic/go-elasticsearch/v8/benchmarks"
	"github.com/elastic/go-elasticsearch/v8/benchmarks/runner"
)

func init() {
	benchmarks.Register(
		benchmarks.Action{
			Name:           "index",
			Category:       "core",
			NumWarmups:     100,
			NumRepetitions: 10000,
			SetupFunc: func(n int, c runner.Config) (*esapi.Response, error) {
				var (
					res *esapi.Response
					err error

					indexName = "test-bench-index"
				)
				res, _ = c.RunnerClient.Indices.Delete([]string{indexName})
				if res != nil && res.Body != nil {
					io.Copy(ioutil.Discard, res.Body)
					res.Body.Close()
				}
				res, err = c.RunnerClient.Indices.Create(indexName)
				if err != nil {
					return res, err
				}
				if res.Body != nil {
					io.Copy(ioutil.Discard, res.Body)
					res.Body.Close()
				}
				return res, err
			},
			RunnerFunc: func(n int, c runner.Config) (*esapi.Response, error) {
				var (
					res *esapi.Response
					err error

					indexName = "test-bench-index"
				)

				docID := fmt.Sprintf("%04d-%04d", n, rand.Intn(benchmarks.DefaultRepetitions))
				docBody := bytes.NewBuffer(benchmarks.DataSources["small"].Bytes())
				res, err = c.RunnerClient.Index(indexName, docBody, c.RunnerClient.Index.WithDocumentID(docID))
				if err != nil {
					return res, err
				}

				var b bytes.Buffer
				if _, err := b.ReadFrom(res.Body); err != nil {
					return nil, err
				}
				res.Body.Close()

				output := gjson.GetBytes(b.Bytes(), "result")
				if output.Str != "created" {
					return nil, fmt.Errorf("Unexpected output: %s", b.String())
				}
				return res, err
			},
		},
	)
}
