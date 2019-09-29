package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"kong-prometheus-exporter/libs"
	"strconv"
)

type Exporter struct {
	//gauge    prometheus.Gauge
	gaugeVec prometheus.GaugeVec
}

func NewExporter(metricsPrefix string) *Exporter {
	//gauge := prometheus.NewGauge(prometheus.GaugeOpts{
	//	Namespace: metricsPrefix,
	//	Name:      "http_status",
	//	Help:      "This is a kong http_status metric"})

	gaugeVec := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "http_status",
		Help:      "This is a kong http_status Vec metric"},
		[]string{"label"})

	return &Exporter{
		//gauge:    gauge,
		gaugeVec: gaugeVec,
	}
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	var metricsKey string
	var metricsValue string
	var metricsValueFloat64 float64
	var metricsCount float64
	var url string
	//var  a string
	ips := libs.GetKongPodIP("ops")
	for ip := ips.Front(); ip != nil; ip = ip.Next() {
		//fmt.Print(ip.Value)
		//fmt.Printf("%T\n",ip.Value)
		url = "http://" + ip.Value.(string) + ":8001/metrics"
		//url = "http://52.184.29.61:8001/metrics"
		metricsKey, metricsValue = GetMetrics(url)
		//metrics_value += metrics_value
		metricsValueFloat64, _ = strconv.ParseFloat(metricsValue, 64)
		fmt.Printf("%s metrics:%g \n", url, metricsValueFloat64)
		metricsCount += metricsValueFloat64
	}


	//e.gauge.Set(metrics_value_float64)

	e.gaugeVec.WithLabelValues(metricsKey).Set(metricsCount)
	//e.gauge.Collect(ch)
	e.gaugeVec.Collect(ch)
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	//e.gauge.Describe(ch)
	e.gaugeVec.Describe(ch)
}
