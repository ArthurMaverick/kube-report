package client

import (
	"context"
	"log"

	v1App "k8s.io/api/apps/v1"
	v1Core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

var (
	hasMetrics bool = false
)

type IKubernetesClient interface {
	GetAllDeploymentsPerNamespaces(namespaces []v1Core.Namespace) []*v1App.DeploymentList
	GetAllDeploymentsPerNamespace(namespaces string) *v1App.DeploymentList
	GetDeployment(namespace, deployment string) *v1App.Deployment
	ListNamespaces() []v1Core.Namespace
	PodMetricsList(namespace string) *v1beta1.PodMetricsList
}

func (k *K8sClient) GetAllDeploymentsPerNamespaces(namespaces []v1Core.Namespace) []*v1App.DeploymentList {
	deployApi := k.coreClient.AppsV1()

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
	deployApi := k.coreClient.AppsV1()

	deploy := deployApi.Deployments(namespaces)
	list, err := deploy.List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		log.Fatalf("Error listing deployment: %v", err)
	}

	return list
}

func (k *K8sClient) GetDeployment(namespace, deployment string) *v1App.Deployment {
	deployApi := k.coreClient.AppsV1()

	deploy := deployApi.Deployments(namespace)
	getDeploy, err := deploy.Get(context.TODO(), deployment, metav1.GetOptions{})

	if err != nil {
		log.Fatalf("Error listing deployment: %v", err)
	}

	return getDeploy
}

func (k *K8sClient) ListNamespaces() []v1Core.Namespace {
	ns := k.coreClient.CoreV1().Namespaces()

	listNamespaces, err := ns.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing namespaces: %v", err)
	}

	return listNamespaces.Items
}

func (k K8sClient) PodMetricsList(namespace string) *v1beta1.PodMetricsList {
	if !hasMetrics {
		return nil
	}

	podMetrics := k.metrisClient.MetricsV1beta1().PodMetricses("")

	podMetricsList, err := podMetrics.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing pod metrics: %v", err)
	}

	return podMetricsList
}

// func (k *K8sClient) DeploymentMetrics(namespace, deployment string) *v1beta1.PodMetricsList {
// 	deploy := k.GetDeployment(namespace, deployment)
// 	podMetrics := k.PodMetricsList(namespace)

// 	var deploymentMetrics v1beta1.PodMetricsList
// 	for _, pod := range podMetrics.Items {
// 		for _, container := range deploy.Spec.Template.Spec.Containers {
// 			for _, containerMetric := range pod.Containers {
// 				if container.Name == containerMetric.Name {
// 					deploymentMetrics.Items = append(deploymentMetrics.Items, pod)
// 				}
// 			}
// 		}
// 	}
// 	return &deploymentMetrics
// }
