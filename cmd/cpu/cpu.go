package cpu

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
)

func PrintCpuInfo() {
	cpuInfo, _ := GetCpuInfo()
	percent, _ := cpu.Percent(120, false)
	// number of cores
	numberOfCores := len(cpuInfo)

	fmt.Printf("Model:           %s\n", cpuInfo[0].ModelName)
	fmt.Printf("Number of cores: %d\n", numberOfCores)
	fmt.Printf("CPU Usage:       %f%%\n", percent[0])
}

func GetCpuInfo() ([]cpu.InfoStat, error) {
	return cpu.Info()
}
