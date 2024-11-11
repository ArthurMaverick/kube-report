package main

import (
	"fmt"

	k8s "github.com/ArthurMaverick/kube-report/kubernetes"
)

// nodes:
// Pegar versao do cluster
// numero de nodes com informacoes de systema operacional, memoria, cpu de cada node
// distribuicao eks, gke, aks e onprem
// numero de nodes por versao do k8s
func main() {
	c := k8s.NewK8sClient()
	response := c.GetAllDeploymentsPerNamespaces(c.ListNamespaces())
	for _, list := range response {
		for _, item := range list.Items {
			fmt.Println("--- DEPLOY INFO ---")
			fmt.Printf("Namespace: %s\n", item.Namespace)
			fmt.Printf("Name: %s\n", item.Name)
			fmt.Printf("Replicas: %d\n", *item.Spec.Replicas)
			fmt.Printf("Containers: %d\n", len(item.Spec.Template.Spec.Containers))
			fmt.Printf("Image: %s\n", item.Spec.Template.Spec.Containers[0].Image)
			fmt.Println("--- RESOURCES ---")
			fmt.Printf("CPURequest: %s\n", item.Spec.Template.Spec.Containers[0].Resources.Requests.Cpu().String())
			fmt.Printf("CPULimit: %s\n", item.Spec.Template.Spec.Containers[0].Resources.Limits.Cpu().String())
			fmt.Printf("MemoryRequest: %s\n", item.Spec.Template.Spec.Containers[0].Resources.Requests.Memory().String())
			fmt.Printf("MemoryLimit: %s\n", item.Spec.Template.Spec.Containers[0].Resources.Limits.Memory().String())
			fmt.Println("--- IMAGE ---")
			fmt.Println("================================")
		}
	}
}
