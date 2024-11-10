package kubernetes

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	v1App "k8s.io/api/apps/v1"
	v1Core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func (k *K8sClient) ListAllDeployments(namespaces []v1Core.Namespace) []*v1App.DeploymentList {
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

func (k *K8sClient) ListDeploymentsPerNamespace(namespaces string) *v1App.DeploymentList {
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

// ApplyYAML applies a Kubernetes manifest from a YAML file
func (k *K8sClient) ApplyYAML(manifestPath string) error {
	// Open the YAML file
	file, err := os.Open(manifestPath)
	if err != nil {
		return fmt.Errorf("failed to open manifest file: %w", err)
	}
	defer file.Close()

	// Create a YAML decoder for multi-document YAML
	decoder := yaml.NewYAMLOrJSONDecoder(file, 4096)

	for {
		// Decode each document into an Unstructured object
		obj := &unstructured.Unstructured{}
		if err := decoder.Decode(obj); err != nil {
			if err.Error() == "EOF" {
				break // End of file
			}
			return fmt.Errorf("failed to decode YAML: %w", err)
		}

		// Get namespace from the object or set it to 'default'
		namespace := obj.GetNamespace()
		if namespace == "" {
			namespace = "default"
		}
		obj.SetNamespace(namespace)

		// Get the GroupVersionResource (GVR)
		gvk := obj.GroupVersionKind()
		gvr := schema.GroupVersionResource{
			Group:    gvk.Group,
			Version:  gvk.Version,
			Resource: resourceName(gvk.Kind),
		}

		// Use the dynamic client to get the existing resource
		resourceClient := k.dyn.Resource(gvr).Namespace(namespace)
		existing, err := resourceClient.Get(context.TODO(), obj.GetName(), metav1.GetOptions{})

		if errors.IsNotFound(err) {
			// Create the resource if it doesn't exist
			_, err = resourceClient.Create(context.TODO(), obj, metav1.CreateOptions{})
			if err != nil {
				return fmt.Errorf("failed to create resource: %w", err)
			}
			fmt.Printf("Resource %s created in namespace %s\n", obj.GetName(), namespace)
		} else if err == nil {
			// Set the resourceVersion from the existing object
			obj.SetResourceVersion(existing.GetResourceVersion())

			// Update the resource
			_, err = resourceClient.Update(context.TODO(), obj, metav1.UpdateOptions{})
			if err != nil {
				return fmt.Errorf("failed to update resource: %w", err)
			}
			fmt.Printf("Resource %s updated in namespace %s\n", obj.GetName(), namespace)
		} else {
			return fmt.Errorf("failed to get resource: %w", err)
		}
	}

	return nil
}

// resourceName converts a Kind to its plural, lowercase form
func resourceName(kind string) string {
	return strings.ToLower(kind) + "s" // Basic pluralization
}
