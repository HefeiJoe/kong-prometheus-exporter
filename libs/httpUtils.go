package libs

import (
	"io/ioutil"
	"net/http"
	"github.com/wonderivan/logger"
)

func Get(url string) (response string) {
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
