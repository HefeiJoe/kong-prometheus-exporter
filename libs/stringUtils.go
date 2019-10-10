package libs

import (
	"encoding/json"
	"github.com/wonderivan/logger"
	"strings"
)

func StrToJsonStr(str string) (jsonStr string) {
	if str == "" {
		logger.Error("str is null!")
	}
	return strings.ReplaceAll(strings.ReplaceAll("{\""+strings.Split(str,"{")[1],"=","\":"),",",",\"")
}
func StrToMap(str string) (dat map[string]string) {
	if err := json.Unmarshal([]byte(StrToJsonStr(str)), &dat); err != nil {
		logger.Error(err)
	}
	return dat
}
