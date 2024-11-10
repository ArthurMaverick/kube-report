package kubernetes

import (
	"flag"
	"log"

	// "k8s.io/client-go/kubernetes"
	"path/filepath"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type K8sClient struct {
	client *kubernetes.Clientset
	dyn    dynamic.Interface
}

func NewK8sClient() *K8sClient {
	var kubeconfig *string
	var config *rest.Config

	config, err := rest.InClusterConfig()

	if err != nil {
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "optional - absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()

		localConfig, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

		if err != nil {
			panic(err.Error())
		}
		config = localConfig
	}

	kubernetesClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}

	return &K8sClient{
		client: kubernetesClient,
		dyn:    dynamicClient,
	}
}
