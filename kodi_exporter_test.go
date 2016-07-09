// Copyright (C) 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	logrus "github.com/Sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

type kodiserver struct {
	*httptest.Server
}

func newKodiServer(resp string) *kodiserver {
	h := &kodiserver{}
	h.Server = httptest.NewServer(handler(h, resp))
	return h
}

func handler(ks *kodiserver, resp string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}
}

func TestKodiExporter(t *testing.T) {
	h := newKodiServer(`{"id":1,"jsonrpc":"2.0","result":"pong"}`)
	defer h.Close()

	collector, err := NewExporter(h.URL, "", "")
	if err != nil {
		t.Fatalf("%v", err)
	}

	ch := make(chan prometheus.Metric)
	go func() {
		defer close(ch)
		collector.Collect(ch)
	}()
}
