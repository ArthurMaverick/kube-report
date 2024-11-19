package utils

import (
	"fmt"
	"strconv"
)

func FormatCpu(cpuUsage string) string {
	if cpuUsage == "1" {
		return "1m"
	}
	return cpuUsage
}

func FormatMemory(memUsage string) string {
	value, err := strconv.ParseInt(memUsage, 10, 64)
	if err == nil {
		if value >= 1024*1024*1024 {
			return fmt.Sprintf("%.2fGB", float64(value)/1024/1024/1024)
		} else if value >= 1024*1024 {
			return fmt.Sprintf("%.2fMB", float64(value)/1024/1024)
		}
	}
	return memUsage
}
