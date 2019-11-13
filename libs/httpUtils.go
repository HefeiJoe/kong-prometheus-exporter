package libs

import (
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/json"
	_struct "kong-prometheus-exporter/struct"
	"net/http"
	"github.com/wonderivan/logger"
)

func GetStruct(url string) (*_struct.AllData ) {
	var allData  = _struct.AllData{}
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		logger.Error(err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
		return nil
	}
	err = json.Unmarshal(body, &allData)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return &allData
}

func GetStr(url string) (response string) {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		logger.Error(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
		return ""
	}
	return string(body)
}
