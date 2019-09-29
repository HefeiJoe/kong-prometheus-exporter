package collector

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)


func Get(url string) (response string) {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

func GetMetrics(url string)  (string, string) {
	var metrics_key string
	var metrics_value string
	metricsResponse :=Get(url)
	if metricsResponse != "" {
		//fmt.Print(metricsResponse)
		flysnowRegexp := regexp.MustCompile(`kong_http_status.*\d`)
		params := flysnowRegexp.FindStringSubmatch(metricsResponse)
		//fmt.Print(params)

		for _,param :=range params {
			//fmt.Println(param)
			metrics := strings.Split(param, " ")
			//fmt.Print(metrics)
			metrics_key = metrics[0]
			metrics_value = metrics[1]
		}

	}
	//fmt.Print(metrics_key)
	//fmt.Print(metrics_value)
	//fmt.Print(metricsResponse)
	return metrics_key, metrics_value
}