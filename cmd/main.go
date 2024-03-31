package main

import (
	"fmt"

	"github.com/berubenic/resource-monitor/cmd/cpu"
	"github.com/berubenic/resource-monitor/cmd/memory"
)

func main() {
	fmt.Println("****** System Info *******")
	fmt.Println()
	fmt.Println("***** Virtual Memory *****")
	memory.PrintVirtualMemory()
	fmt.Println()
	fmt.Println("******** CPU Info ********")
	cpu.PrintCpuInfo()
}
