package configs

import (
	"os"
	"github.com/wonderivan/logger"
       )

func InitConfig() Conf {
	var c Conf
	c.getConfigFromEnv()
	return c
}

type Conf struct {
	serviceName     string
	namespace string
	port string
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
	c.serviceName=serviceName
	c.namespace = namespace
	c.port = port
	return c
}