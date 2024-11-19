package main

import (
	"github.com/ArthurMaverick/kube-report/pkg/clients"
	"github.com/ArthurMaverick/kube-report/pkg/data"
	"github.com/ArthurMaverick/kube-report/pkg/service"
)

func main() {
	k8sClient := clients.NewK8sClient()
	fmtClient := data.NewFormatters(k8sClient)
	outputClient := service.NewJsonOutput(fmtClient)
	outputClient.JsonFile()
}
