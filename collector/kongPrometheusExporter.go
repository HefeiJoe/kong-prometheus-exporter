package collector

import (
	"github.com/wonderivan/logger"
	"github.com/prometheus/client_golang/prometheus"
	"kong-prometheus-exporter/libs"
	"os"
	"strconv"
	"strings"
)
const (
	BANDWIDTH            = "kong_bandwidth"
	DATASTORE_REACHABLE      = "kong_datastore_reachable"
	HTTP_STATUS = "kong_http_status"
	LATENCY_BUCKET = "kong_latency_bucket"
	LATENCY_COUNT = "kong_latency_count"
	LATENCY_SUM = "kong_latency_sum"
	NGINX_HTTP_CURRENT_CONNECTIONS = "kong_nginx_http_current_connections"
	NGINX_METRIC_ERRORS_TOTAL = "kong_nginx_metric_errors_total"

)
type Exporter struct {
	bandwidth prometheus.GaugeVec
	datastore_reachable prometheus.Gauge
	http_status prometheus.GaugeVec
	latency_bucket prometheus.GaugeVec
	latency_count prometheus.GaugeVec
	latency_sum prometheus.GaugeVec
	nginx_http_current_connections prometheus.GaugeVec
	nginx_metric_errors_total prometheus.Gauge
}

func NewExporter(metricsPrefix string) *Exporter {
	bandwidth := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "bandwidth",
		Help:      "This is a kong bandwidth metric"},
		[]string{
			"type",
			"service",
		})
	datastore_reachable := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "datastore_reachable",
		Help:      "This is a kong datastore_reachable metric"})
	http_status := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "http_status",
		Help:      "This is a kong http_status metric"},
		[]string{
			"code",
			"service",
		})
	latency_bucket := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "latency_bucket",
		Help:      "This is a kong bandwidth metric"},
		[]string{
			"type",
			"service",
			"le",
		})
	latency_count := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "latency_count",
		Help:      "This is a kong latency_count metric"},
		[]string{
			"type",
			"service",
		})
	latency_sum := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "latency_sum",
		Help:      "This is a kong latency_sum metric"},
		[]string{
			"type",
			"service",
		})
	nginx_http_current_connections := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "nginx_http_current_connections",
		Help:      "This is a kong nginx_http_current_connections metric"},
		[]string{"state"})
	nginx_metric_errors_total := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "nginx_metric_errors_total",
		Help:      "This is a kong nginx_metric_errors_total metric"})

	return &Exporter{
		bandwidth: bandwidth,
		datastore_reachable: datastore_reachable,
		http_status: http_status,
		latency_bucket:latency_bucket,
		latency_count:latency_count,
		latency_sum:latency_sum,
		nginx_http_current_connections:nginx_http_current_connections,
		nginx_metric_errors_total:nginx_metric_errors_total,
	}
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	var url string
	var response string
	var newResponse string
	namespace:=os.Getenv("NAMESPACE")
	port:=os.Getenv("PORT")
	if namespace == "" {
		logger.Error("Namespace cannot be empty!")
	}
	if port == "" {
		logger.Error("Port cannot be empty!")
	}
	ips := libs.GetKongPodIP(namespace)
	if ips.Len()==0{
		logger.Error("ips is null!")
	}
	for ip := ips.Front(); ip != nil; ip = ip.Next() {
		url = "http://" + ip.Value.(string) + ":"+port+"/metrics"
		metricsResponse :=Get(url)
		logger.Info(url)
		response = response+ metricsResponse
	}
	logger.Info("response:",response)
	if response != "" {
		metricsList :=strings.Split(response,"\n")
		for i := 0; i < len(metricsList)-1; i++ {
			k:=strings.Split(metricsList[i]," ")[0]
			var v float64
			var metricsValueFloat64 float64
			if k==DATASTORE_REACHABLE{
				metricsValueFloat64, _ = strconv.ParseFloat(strings.Split(metricsList[i]," ")[1], 64)
				newResponse = newResponse + k + " " + strconv.FormatFloat(metricsValueFloat64,'E',-1,64) +"\n"
				e.datastore_reachable.Set(metricsValueFloat64)
			}else {
				if strings.HasPrefix(k, "#") {
					if !strings.Contains(newResponse, metricsList[i]) {
						newResponse = newResponse + metricsList[i]+"\n"
					}
				} else {
					if !strings.Contains(newResponse, k) {
						for j := i; j < len(metricsList)-1; j++ {
							kv := strings.Split(metricsList[j], " ")
							if kv[0] == k {
								metricsValueFloat64, _ = strconv.ParseFloat(kv[1], 64)
								v = v + metricsValueFloat64
							}
						}
						newResponse = newResponse + k + " " + strconv.FormatFloat(v, 'E', -1, 64) + "\n"
						kMap := StrToMap(k)
						if strings.Contains(k,BANDWIDTH){
							e.bandwidth.WithLabelValues(kMap["type"], kMap["service"]).Set(v)
						}
						if strings.Contains(k,HTTP_STATUS){
							e.http_status.WithLabelValues(kMap["code"], kMap["service"]).Set(v)
						}
						if strings.Contains(k,LATENCY_BUCKET){
							e.latency_bucket.WithLabelValues(kMap["type"], kMap["service"], kMap["le"]).Set(v)
						}
						if strings.Contains(k,LATENCY_COUNT){
							e.latency_count.WithLabelValues(kMap["type"], kMap["service"]).Set(v)
						}
						if strings.Contains(k,LATENCY_SUM){
							e.latency_sum.WithLabelValues(kMap["type"], kMap["service"]).Set(v)
						}
						if strings.Contains(k,NGINX_HTTP_CURRENT_CONNECTIONS){
							e.nginx_http_current_connections.WithLabelValues(kMap["state"]).Set(v)
						}
						if strings.Contains(k,NGINX_METRIC_ERRORS_TOTAL){
							e.nginx_metric_errors_total.Set(v)
						}
					}
				}
			}
		}
	}
	e.bandwidth.Collect(ch)
	e.datastore_reachable.Collect(ch)
	e.http_status.Collect(ch)
	e.latency_bucket.Collect(ch)
	e.latency_count.Collect(ch)
	e.latency_sum.Collect(ch)
	e.nginx_http_current_connections.Collect(ch)
	e.nginx_metric_errors_total.Collect(ch)

}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.bandwidth.Describe(ch)
	e.datastore_reachable.Describe(ch)
	e.http_status.Describe(ch)
	e.latency_bucket.Describe(ch)
	e.latency_count.Describe(ch)
	e.latency_sum.Describe(ch)
	e.nginx_http_current_connections.Describe(ch)
	e.nginx_metric_errors_total.Describe(ch)
}
