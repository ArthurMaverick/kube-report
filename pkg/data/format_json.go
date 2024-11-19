package formatters

import (
	"log"

	"github.com/ArthurMaverick/kube-report/pkg/utils"
	v1 "k8s.io/api/core/v1"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

type DeployInfo struct {
	Namespace  string      `json:"namespace"`
	Name       string      `json:"name"`
	Replicas   int32       `json:"replicas"`
	Containers []Container `json:"containers"`
}

type Container struct {
	Image      string `json:"image"`
	CpuRequest string `json:"cpu_request"`
	CpuLimit   string `json:"cpu_limit"`
	MemRequest string `json:"mem_request"`
	MemLimit   string `json:"mem_limit"`
	CurrentCpu string `json:"current_cpu"`
	CurrentMem string `json:"current_mem"`
}

func (d *Formats) FormatJSONData() *[]DeployInfo {
	var deployInfos []DeployInfo
	deploy := d.k8sClient.GetAllDeploymentsPerNamespaces(d.k8sClient.ListNamespaces())

	for _, list := range deploy {
		if len(list.Items) == 0 {
			continue
		}

		podMetrics := d.k8sClient.PodMetricsList(list.Items[0].Namespace)

		for _, item := range list.Items {
			deployInfo := DeployInfo{
				Namespace:  item.Namespace,
				Name:       item.Name,
				Replicas:   *item.Spec.Replicas,
				Containers: getAllContainersInfo(item.Spec.Template.Spec.Containers, podMetrics),
			}
			deployInfos = append(deployInfos, deployInfo)
		}
	}
	return &deployInfos
}

func getAllContainersInfo(container []v1.Container, podMetrics *v1beta1.PodMetricsList) []Container {
	var containersInfo []Container
	for _, c := range container {
		containerInfo := Container{
			Image:      c.Image,
			CpuRequest: c.Resources.Requests.Cpu().String(),
			CpuLimit:   c.Resources.Limits.Cpu().String(),
			MemRequest: c.Resources.Requests.Memory().String(),
			MemLimit:   c.Resources.Limits.Memory().String(),
		}

		if podMetrics != nil {
			for _, metricsContainer := range podMetrics.Items {
				for _, containerMetrics := range metricsContainer.Containers {
					if containerMetrics.Name == c.Name {
						// todo: refactor this
						rawCpu := containerMetrics.Usage.Cpu().MilliValue()
						rawMem := containerMetrics.Usage.Memory().MilliValue()
						log.Printf("Container: %s, Raw CPU: %v, Raw Memory: %v\n", containerMetrics.Name, rawCpu, rawMem)

						containerInfo.CurrentCpu = utils.FormatCpu(containerMetrics.Usage.Cpu().String())
						containerInfo.CurrentMem = utils.FormatMemory(containerMetrics.Usage.Memory().String())
					}
				}
			}
		}

		containersInfo = append(containersInfo, containerInfo)
	}
	return containersInfo
}
