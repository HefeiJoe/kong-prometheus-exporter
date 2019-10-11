package configs

import (
	"os"
	"github.com/wonderivan/logger"
)
var Cf *Conf

func InitConfig() {
	var c Conf
	Cf = c.getConfigFromEnv()
}

type Conf struct {
	ServiceName     string
	Namespace string
	Port string
}


func (c *Conf) getConfigFromEnv() *Conf {
	serviceName := os.Getenv("SERVICE_NAME")
	namespace := os.Getenv("NAMESPACE")
	port:= os.Getenv("PORT")
	if serviceName == "" {
		logger.Error("namespace is empty")
	}
	if namespace == "" {
		logger.Error("namespace is empty")
	}
	if port == "" {
		logger.Error("port is empty")
	}
	cf := &Conf{}
	cf.ServiceName=serviceName
	cf.Namespace = namespace
	cf.Port = port
	return c
}
