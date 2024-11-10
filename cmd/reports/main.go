package main

import (
	"fmt"

	k8s "github.com/ArthurMaverick/kube-report/kubernetes"
)

func main() {
	c := k8s.NewK8sClient()
	fmt.Println(c.ListDeploymentsPerNamespace("default"))
	// fmt.Println(c.ListAllDeployments(c.ListNamespaces()))
}
