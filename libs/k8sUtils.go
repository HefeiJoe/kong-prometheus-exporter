package libs

import (
	"container/list"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var (
	c *kubernetes.Clientset
)

func InitK8sClient() {

	//从集群中读取kubectl的config信息
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	c, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	//从外部读取kubectl的config信息
	//config, err := clientcmd.BuildConfigFromFlags("", *Kubeconfig)
	//if err != nil {
	//	panic(err.Error())
	//}
	//c, err = kubernetes.NewForConfig(config)
	//if err != nil {
	//	panic(err.Error())
	//}
}


func GetKongPodIP(namespace string) (ips list.List) {
	//namespace := "ops"
	var ipList list.List
 	endpoints, _ := c.CoreV1().Endpoints(namespace).Get("kong3-kong-admin", metav1.GetOptions{})
	for _, endpoint := range endpoints.Subsets {
		for _, ip := range endpoint.Addresses {
			ipList.PushFront(ip.IP)
		}

	}
	//for i := ipList.Front(); i != nil; i = i.Next() {
	//	fmt.Print(i.Value, " ")
	//}
	//for i := adminPortList.Front(); i != nil; i = i.Next() {
	//	fmt.Print(i.Value, " ")
	//}
	return ipList

}