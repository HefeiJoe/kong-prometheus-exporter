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
	//从外部读取kubectl的config信息
	/*config, err := clientcmd.BuildConfigFromFlags("", *Kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	c, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}*/
}


func GetKongPodIP(namespace string) (ips list.List) {
	var ipList list.List
	kongName:=os.Getenv("KONG_NAME")
	if kongName == "" {
		logger.Error("kongName cannot be empty!")
	}
	logger.Info("kongName: ",kongName)
 	endpoints, err := c.CoreV1().Endpoints(namespace).Get(kongName, metav1.GetOptions{})
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