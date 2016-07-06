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
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	prom_version "github.com/prometheus/common/version"

	"github.com/nlamirault/kodi_exporter/kodi"
	"github.com/nlamirault/kodi_exporter/version"
)

const (
	namespace = "kodi"
)

var (
	up = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "up"),
		"Was the last query of Kodi successful.",
		nil, nil,
	)
)

// Exporter collects Consul stats from the given server and exports them using
// the prometheus metrics package.
type Exporter struct {
	URI    string
	Client *kodi.Client
}

// NewExporter returns an initialized Exporter.
func NewExporter(uri string, username string, password string) (*Exporter, error) {
	// Set up our Kodi client connection.
	log.Infoln("Setup Kodi client")
	client := kodi.NewClient(uri, username, password)

	params := map[string]interface{}{
		`title`:   `Prometheus`,
		`message`: `Prometheus exporter for Kodi is ready`,
	}
	resp := &kodi.ShowNotificationResponse{}
	err := client.RPC(&kodi.Request{
		Jsonrpc: "2.0",
		Method:  "GUI.ShowNotification",
		Id:      1,
		Params:  params}, resp)
	if err != nil {
		return nil, err
	}
	log.Infof("Kodi connection: %s\n", resp.Result)
	log.Infoln("Init exporter")
	// Init our exporter.
	return &Exporter{
		URI:    uri,
		Client: client,
	}, nil
}

// Describe describes all the metrics ever exported by the Kodi exporter.
// It implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- up
	// ch <- clusterServers
	// ch <- nodeCount
	// ch <- serviceCount
	// ch <- serviceNodesHealthy
	// ch <- nodeChecks
	// ch <- serviceChecks
	// ch <- keyValues
}

// Collect fetches the stats from configured Consul location and delivers them
// as Prometheus metrics.
// It implements prometheus.Collector.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {

	// 	// How many peers are in the Consul cluster?
	// 	peers, err := e.client.Status().Peers()
	// 	if err != nil {
	// 		ch <- prometheus.MustNewConstMetric(
	// 			up, prometheus.GaugeValue, 0,
	// 		)
	// 		log.Errorf("Query error is %v", err)
	// 		return
	// 	}

	// 	// We'll use peers to decide that we're up.
	// 	ch <- prometheus.MustNewConstMetric(
	// 		up, prometheus.GaugeValue, 1,
	// 	)
	// 	ch <- prometheus.MustNewConstMetric(
	// 		clusterServers, prometheus.GaugeValue, float64(len(peers)),
	// 	)

	// 	// How many nodes are registered?
	// 	nodes, _, err := e.client.Catalog().Nodes(&consul_api.QueryOptions{})
	// 	if err != nil {
	// 		// FIXME: How should we handle a partial failure like this?
	// 	} else {
	// 		ch <- prometheus.MustNewConstMetric(
	// 			nodeCount, prometheus.GaugeValue, float64(len(nodes)),
	// 		)
	// 	}

	// 	// Query for the full list of services.
	// 	serviceNames, _, err := e.client.Catalog().Services(&consul_api.QueryOptions{})
	// 	if err != nil {
	// 		// FIXME: How should we handle a partial failure like this?
	// 		return
	// 	}
	// 	ch <- prometheus.MustNewConstMetric(
	// 		serviceCount, prometheus.GaugeValue, float64(len(serviceNames)),
	// 	)

	// 	if e.healthSummary {
	// 		e.collectHealthSummary(ch, serviceNames)
	// 	}

	// 	checks, _, err := e.client.Health().State("any", &consul_api.QueryOptions{})
	// 	if err != nil {
	// 		log.Errorf("Failed to query service health: %v", err)
	// 		return
	// 	}

	// 	for _, hc := range checks {
	// 		var passing float64
	// 		if hc.Status == consul.HealthPassing {
	// 			passing = 1
	// 		}
	// 		if hc.ServiceID == "" {
	// 			ch <- prometheus.MustNewConstMetric(
	// 				nodeChecks, prometheus.GaugeValue, passing, hc.CheckID, hc.Node,
	// 			)
	// 		} else {
	// 			ch <- prometheus.MustNewConstMetric(
	// 				serviceChecks, prometheus.GaugeValue, passing, hc.CheckID, hc.Node, hc.ServiceID,
	// 			)
	// 		}
	// 	}

	// 	e.collectKeyValues(ch)
}

// collectHealthSummary collects health information about every node+service
// combination. It will cause one lookup query per service.
// func (e *Exporter) collectHealthSummary(ch chan<- prometheus.Metric, serviceNames map[string][]string) {
// 	for s := range serviceNames {
// 		service, _, err := e.client.Health().Service(s, "", false, &consul_api.QueryOptions{})
// 		if err != nil {
// 			log.Errorf("Failed to query service health: %v", err)
// 			continue
// 		}

// 		for _, entry := range service {
// 			// We have a Node, a Service, and one or more Checks. Our
// 			// service-node combo is passing if all checks have a `status`
// 			// of "passing."
// 			passing := 1.
// 			for _, hc := range entry.Checks {
// 				if hc.Status != consul.HealthPassing {
// 					passing = 0
// 					break
// 				}
// 			}
// 			ch <- prometheus.MustNewConstMetric(
// 				serviceNodesHealthy, prometheus.GaugeValue, passing, entry.Service.ID, entry.Node.Node,
// 			)
// 		}
// 	}
// }

// func (e *Exporter) collectKeyValues(ch chan<- prometheus.Metric) {
// 	if e.kvPrefix == "" {
// 		return
// 	}

// 	kv := e.client.KV()
// 	pairs, _, err := kv.List(e.kvPrefix, &consul_api.QueryOptions{})
// 	if err != nil {
// 		log.Errorf("Error fetching key/values: %s", err)
// 		return
// 	}

// 	for _, pair := range pairs {
// 		if e.kvFilter.MatchString(pair.Key) {
// 			val, err := strconv.ParseFloat(string(pair.Value), 64)
// 			if err == nil {
// 				ch <- prometheus.MustNewConstMetric(
// 					keyValues, prometheus.GaugeValue, val, pair.Key,
// 				)
// 			}
// 		}
// 	}
// }

// func init() {
// 	prometheus.MustRegister(prom_version.NewCollector("kodi_exporter"))
// }

func main() {
	var (
		showVersion   = flag.Bool("version", false, "Print version information.")
		listenAddress = flag.String("web.listen-address", ":9111", "Address to listen on for web interface and telemetry.")
		metricsPath   = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
		kodiServer    = flag.String("kodi.server", "localhost:9090", "HTTP API address of the Kodi server.")
		kodiUsername  = flag.String("kodi.username", "", "Username for authentication to the Kodi server.")
		kodiPassword  = flag.String("kodi.password", "", "Password for authentication to the Kodi server.")
	)
	flag.Parse()

	if *showVersion {
		fmt.Printf("Kodi Prometheus exporter. v%s\n", version.Version)
		// fmt.Fprintln(os.Stdout, version.Print("kodi_exporter"))
		os.Exit(0)
	}

	log.Infoln("Starting kodi_exporter", prom_version.Info())
	// log.Infoln("Build context", prom_version.BuildContext())

	exporter, err := NewExporter(*kodiServer, *kodiUsername, *kodiPassword)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	log.Infoln("Register exporter")
	prometheus.MustRegister(exporter)

	http.Handle(*metricsPath, prometheus.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Kodi Exporter</title></head>
             <body>
             <h1>Kodi Exporter</h1>
             <p><a href='` + *metricsPath + `'>Metrics</a></p>
             </body>
             </html>`))
	})

	log.Infoln("Listening on", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
