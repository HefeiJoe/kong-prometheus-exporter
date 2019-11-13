package configs

import (
	"github.com/wonderivan/logger"
	"os"
)
var Cf *Conf

func InitConfig() {
	var c Conf
	Cf = c.getConfigFromEnv()
}

type Conf struct {
	ServiceName string
	Namespace string
	KongAdminPort string
	Port string
}


func (c *Conf) getConfigFromEnv() *Conf {
	serviceName := os.Getenv("SERVICE_NAME")
	namespace := os.Getenv("NAMESPACE")
	kongAdminPort:=os.Getenv("kONG_ADMIN_PORT")
	port:= os.Getenv("PORT")
	if serviceName == "" {
		logger.Error("namespace is empty")
	}
	if namespace == "" {
		logger.Error("namespace is empty")
	}
	if kongAdminPort == "" {
		logger.Error("port is empty")
	}
	if port == "" {
		logger.Error("port is empty")
	}
	cf := &Conf{}
	cf.ServiceName=serviceName
	cf.Namespace = namespace
	cf.KongAdminPort = kongAdminPort
	cf.Port = port
	return cf
}
