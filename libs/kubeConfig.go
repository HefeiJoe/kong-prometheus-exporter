package libs

import (
	"flag"
	"path/filepath"
)
var Kubeconfig *string
func InitConfigConfig() {
	home := "/home/lujin"
	Kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	flag.Parse()
}