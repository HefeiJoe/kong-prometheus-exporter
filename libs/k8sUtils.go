package libs

import (
	"github.com/wonderivan/logger"
	"container/list"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
)

var (
	c *kubernetes.Clientset
)

func InitK8sClient() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	c, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}

func GetKongPodIP(namespace string) (ips list.List) {
	var ipList list.List
	serviceName:=os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		logger.Error("serviceName cannot be empty!")
	}
	logger.Info("kongName: ",serviceName)
 	endpoints, err := c.CoreV1().Endpoints(namespace).Get(serviceName, metav1.GetOptions{})
	if err != nil {
		logger.Error(err)
	}
	for _, endpoint := range endpoints.Subsets {
		for _, ip := range endpoint.Addresses {
			ipList.PushFront(ip.IP)
		}
	}
	return ipList
}