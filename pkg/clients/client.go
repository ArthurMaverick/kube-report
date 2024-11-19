package client

import (
	"flag"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	metricsclient "k8s.io/metrics/pkg/client/clientset/versioned"
)

type K8sClient struct {
	coreClient   *kubernetes.Clientset
	metrisClient *metricsclient.Clientset
}

func NewK8sClient() *K8sClient {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	coreClientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	metricsClientSet, err := metricsclient.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return &K8sClient{
		coreClient:   coreClientSet,
		metrisClient: metricsClientSet,
	}
}
