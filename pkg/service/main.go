package output

import (
	"github.com/ArthurMaverick/kube-report/pkg/data"
)

type JsonOutput struct {
	fmtClient formatters.IFormats
}

func NewJsonOutput(fmtClient formatters.IFormats) *JsonOutput {
	return &JsonOutput{
		fmtClient: fmtClient,
	}
}
