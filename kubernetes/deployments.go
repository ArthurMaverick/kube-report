package kubernetes

import (
	"context"
	"log"

	v1App "k8s.io/api/apps/v1"
	v1Core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *K8sClient) GetAllDeploymentsPerNamespaces(namespaces []v1Core.Namespace) []*v1App.DeploymentList {
	deployApi := k.client.AppsV1()

	var lists []*v1App.DeploymentList

	for _, namespace := range namespaces {
		d := deployApi.Deployments(namespace.Name)
		list, err := d.List(context.TODO(), metav1.ListOptions{})

		if err != nil {
			log.Fatalf("Error listing deployment: %v", err)
		}

		lists = append(lists, list)
	}
	return lists
}

func (k *K8sClient) GetAllDeploymentsPerNamespace(namespaces string) *v1App.DeploymentList {
	deployApi := k.client.AppsV1()

	deploy := deployApi.Deployments(namespaces)
	list, err := deploy.List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		log.Fatalf("Error listing deployment: %v", err)
	}

	return list
}

func (k *K8sClient) GetDeployment(namespace, deployment string) *v1App.Deployment {
	deployApi := k.client.AppsV1()

	deploy := deployApi.Deployments(namespace)
	getDeploy, err := deploy.Get(context.TODO(), deployment, metav1.GetOptions{})

	if err != nil {
		log.Fatalf("Error listing deployment: %v", err)
	}

	return getDeploy
}
