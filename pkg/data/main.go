package formatters

import (
	"github.com/ArthurMaverick/kube-report/pkg/clients"
)

type IFormats interface {
	FormatJSONData() *[]DeployInfo
}

type Formats struct {
	k8sClient client.IKubernetesClient
}

func NewFormatters(k8sClient client.IKubernetesClient) *Formats {
	return &Formats{
		k8sClient: k8sClient,
	}
}
