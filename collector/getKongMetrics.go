package collector

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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

func StrToJsonStr(str string) (jsonStr string) {
	if str == "" {
		Error.Fatalln("str is null!")
	}
	return strings.ReplaceAll(strings.ReplaceAll("{\""+strings.Split(str,"{")[1],"=","\":"),",",",\"")
}
func StrToMap(str string) (dat map[string]string) {
	if err := json.Unmarshal([]byte(StrToJsonStr(str)), &dat); err != nil {
		Error.Fatalln(err)
	}
	return dat
}
