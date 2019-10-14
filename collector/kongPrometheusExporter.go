package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"kong-prometheus-exporter/configs"
	"kong-prometheus-exporter/const"
	"kong-prometheus-exporter/libs"
	"strconv"
	"strings"
)
type Exporter struct {
	bandwidth prometheus.GaugeVec
	datastore_reachable prometheus.GaugeVec
	consumer_http_status prometheus.GaugeVec
	route_http_status prometheus.GaugeVec
	latency_bucket prometheus.GaugeVec
	latency_count prometheus.GaugeVec
	latency_sum prometheus.GaugeVec
	memory_lua_shared_dict_total_bytes prometheus.GaugeVec
	nginx_http_current_connections prometheus.GaugeVec
	nginx_metric_errors_total prometheus.GaugeVec
}

func NewExporter(metricsPrefix string) *Exporter {
	bandwidth := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "bandwidth",
		Help:      "This is a kong bandwidth metric"},
		[]string{
			"type",
			"service",
			"route",
			"podip",
		})
	datastore_reachable := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "datastore_reachable",
		Help:      "This is a kong datastore_reachable metric"},
		[]string{
			"podip",
		})
	route_http_status := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "route_http_status",
		Help:      "This is a kong route_http_status metric"},
		[]string{
			"code",
			"service",
			"route",
			"podip",
		})
	consumer_http_status := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "consumer_http_status",
		Help:      "This is a kong consumer_http_status metric"},
		[]string{
			"code",
			"service",
			"consumer",
			"podip",
		})
	latency_bucket := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "latency_bucket",
		Help:      "This is a kong bandwidth metric"},
		[]string{
			"type",
			"service",
			"route",
			"le",
			"podip",
		})
	latency_count := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "latency_count",
		Help:      "This is a kong latency_count metric"},
		[]string{
			"type",
			"service",
			"route",
			"podip",
		})
	latency_sum := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "latency_sum",
		Help:      "This is a kong latency_sum metric"},
		[]string{
			"type",
			"service",
			"route",
			"podip",
		})
	memory_lua_shared_dict_total_bytes := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "memory_lua_shared_dict_total_bytes",
		Help:      "This is a kong memory_lua_shared_dict_total_bytes metric"},
		[]string{
			"shared_dict",
			"podip",
		})
	nginx_http_current_connections := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "nginx_http_current_connections",
		Help:      "This is a kong nginx_http_current_connections metric"},
		[]string{
			"state",
			"podip",})
	nginx_metric_errors_total := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "nginx_metric_errors_total",
		Help:      "This is a kong nginx_metric_errors_total metric"},
		[]string{
			"podip",
		})

	return &Exporter{
		bandwidth: bandwidth,
		datastore_reachable: datastore_reachable,
		route_http_status: route_http_status,
		consumer_http_status: consumer_http_status,
		latency_bucket:latency_bucket,
		latency_count:latency_count,
		latency_sum:latency_sum,
		memory_lua_shared_dict_total_bytes:memory_lua_shared_dict_total_bytes,
		nginx_http_current_connections:nginx_http_current_connections,
		nginx_metric_errors_total:nginx_metric_errors_total,
	}
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	var url string
	var response string
	resetMetrics(e)
	namespace:=configs.Cf.Namespace
	port:=configs.Cf.Port
	ips := libs.GetKongPodIP(namespace)
	for ip := ips.Front(); ip != nil; ip = ip.Next() {
		url = "http://" + ip.Value.(string) + ":"+port+"/metrics"
		metricsResponse :=libs.Get(url)
		if ips.Len()>1 {
			collectMetrics(metricsResponse, ip.Value.(string), e)
		}
		response = response+ metricsResponse
	}
	collectMetrics(response, "total", e)
	e.bandwidth.Collect(ch)
	e.datastore_reachable.Collect(ch)
	e.route_http_status.Collect(ch)
	e.consumer_http_status.Collect(ch)
	e.latency_bucket.Collect(ch)
	e.latency_count.Collect(ch)
	e.latency_sum.Collect(ch)
	e.memory_lua_shared_dict_total_bytes.Collect(ch)
	e.nginx_http_current_connections.Collect(ch)
	e.nginx_metric_errors_total.Collect(ch)

}

func collectMetrics(response string,mtp string,e *Exporter){
	var newResponse string
	if response != "" {
		metricsList :=strings.Split(response,"\n")
		for i := 0; i < len(metricsList)-1; i++ {
			k:=strings.Split(metricsList[i]," ")[0]
			var v float64
			var metricsValueFloat64 float64
			if k==_const.DATASTORE_REACHABLE{
				metricsValueFloat64, _ = strconv.ParseFloat(strings.Split(metricsList[i]," ")[1], 64)
				newResponse = newResponse + k + " " + strconv.FormatFloat(metricsValueFloat64,'E',-1,64) +"\n"
				e.datastore_reachable.WithLabelValues(mtp).Set(metricsValueFloat64)
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
						kMap := libs.StrToMap(k)
						if strings.Contains(k,_const.BANDWIDTH){
							e.bandwidth.WithLabelValues(kMap["type"], kMap["service"], kMap["route"],mtp).Set(v)
						}
						if strings.Contains(k,_const.ROUTE_HTTP_STATUS){
							e.route_http_status.WithLabelValues(kMap["code"], kMap["service"], kMap["route"],mtp).Set(v)
						}
						if strings.Contains(k,_const.CONSUMER_HTTP_STATUS){
							e.consumer_http_status.WithLabelValues(kMap["code"], kMap["service"], kMap["consumer"],mtp).Set(v)
						}
						if strings.Contains(k,_const.LATENCY_BUCKET){
							e.latency_bucket.WithLabelValues(kMap["type"], kMap["service"], kMap["route"], kMap["le"],mtp).Set(v)
						}
						if strings.Contains(k,_const.LATENCY_COUNT){
							e.latency_count.WithLabelValues(kMap["type"], kMap["service"], kMap["route"],mtp).Set(v)
						}
						if strings.Contains(k,_const.LATENCY_SUM){
							e.latency_sum.WithLabelValues(kMap["type"], kMap["service"], kMap["route"],mtp).Set(v)
						}
						if strings.Contains(k,_const.MEMORY_LUA_SHARED_DICT_TOTAL_BYTES){
							e.memory_lua_shared_dict_total_bytes.WithLabelValues(kMap["shared_dict"],mtp).Set(v)
						}
						if strings.Contains(k,_const.NGINX_HTTP_CURRENT_CONNECTIONS){
							e.nginx_http_current_connections.WithLabelValues(kMap["state"],mtp).Set(v)
						}
						if strings.Contains(k,_const.NGINX_METRIC_ERRORS_TOTAL){
							e.nginx_metric_errors_total.WithLabelValues(mtp).Set(v)
						}
					}
				}
			}
		}
	}
}
func resetMetrics(e *Exporter){
	e.bandwidth.Reset()
	e.datastore_reachable.Reset()
	e.route_http_status.Reset()
	e.consumer_http_status.Reset()
	e.latency_bucket.Reset()
	e.latency_count.Reset()
	e.latency_sum.Reset()
	e.memory_lua_shared_dict_total_bytes.Reset()
	e.nginx_http_current_connections.Reset()
	e.nginx_metric_errors_total.Reset()
}
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.bandwidth.Describe(ch)
	e.datastore_reachable.Describe(ch)
	e.route_http_status.Describe(ch)
	e.consumer_http_status.Describe(ch)
	e.latency_bucket.Describe(ch)
	e.latency_count.Describe(ch)
	e.latency_sum.Describe(ch)
	e.memory_lua_shared_dict_total_bytes.Describe(ch)
	e.nginx_http_current_connections.Describe(ch)
	e.nginx_metric_errors_total.Describe(ch)
}
