package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"kong-prometheus-exporter/collector"
	"kong-prometheus-exporter/configs"
	"kong-prometheus-exporter/libs"
	"net/http"
)

func main() {
	fmt.Println(`
  This is a kong example of prometheus exporter
  Access: http://0.0.0.0:8080
  `)

	//libs.InitConfigConfig()
	libs.InitK8sClient()
	configs.InitConfig()

	// Define parameters
	metricsPath := "/metrics"
	listenAddress := ":8080"
	metricsPrefix := "kong"

	// Register kong exporter, not necessary
	exporter := collector.NewExporter(metricsPrefix)
	prometheus.MustRegister(exporter)

	registry := prometheus.NewRegistry()
	registry.MustRegister(exporter)
	http.Handle(metricsPath, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>A Prometheus kong Exporter</title></head>
			<body>
			<h1>A Prometheus Exporter</h1>
			<p><a href='/metrics'>Metrics</a></p>
			</body>
			</html>`))
	})

	fmt.Println(http.ListenAndServe(listenAddress, nil))
}