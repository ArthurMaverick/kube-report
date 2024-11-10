package kubernetes

import (
	"context"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *K8sClient) ListNamespaces() []v1.Namespace {
	ns := k.client.CoreV1().Namespaces()

	listNamespaces, err := ns.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing namespaces: %v", err)
	}

	// for _, namespace := range list_ns.Items {
	// 	fmt.Println(namespace.Name)
	// }
	return listNamespaces.Items
}
