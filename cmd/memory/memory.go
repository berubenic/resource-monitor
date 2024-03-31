package memory

import (
	"fmt"
	"math"

	"github.com/shirou/gopsutil/v3/mem"
)

func PrintVirtualMemory() {
	v, _ := GetVirtualMemory()

	totalInGB := float64(v.Total) / 1024 / 1024 / 1024
	availableInGB := float64(v.Available) / 1024 / 1024 / 1024

	totalInGB = math.Round(totalInGB*100) / 100
	availableInGB = math.Round(availableInGB*100) / 100

	availablePercent := availableInGB / totalInGB * 100

	// calculat used percent with total and available
	usedInGB := totalInGB - availableInGB
	usedPercent := usedInGB / totalInGB * 100

	fmt.Printf("Total:             %.2f GB\n", totalInGB)
	fmt.Printf("Available:         %.2f GB\n", availableInGB)
	fmt.Printf("Available Percent: %.2f%%\n", availablePercent)
	fmt.Printf("Used:              %.2f GB\n", totalInGB-availableInGB)
	fmt.Printf("Used Percent:      %.2f%%\n", usedPercent)
}

func GetVirtualMemory() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}
