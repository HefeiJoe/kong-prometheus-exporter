package libs

import (
"flag"
"path/filepath"
)
var Kubeconfig *string
func InitConfigConfig() {
	home := "E:/"
	Kubeconfig = flag.String("kubeconfig", filepath.Join(home, "Code", "config"), "(optional) absolute path to the kubeconfig file")
	flag.Parse()
}