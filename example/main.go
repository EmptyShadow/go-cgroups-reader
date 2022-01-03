package main

import (
	"fmt"

	cgroupsreader "github.com/EmptyShadow/go-cgroups-reader/v1"
)

func main() {
	recources, enabled, err := cgroupsreader.ReadResources("/", "")
	if err != nil {
		panic(err)
	}

	fmt.Println("Enabled", enabled)
	fmt.Println("QuotaUS", recources.CPU.QuotaUS)
	fmt.Println("PeriodUS", recources.CPU.PeriodUS)
	fmt.Println("Mem limit", recources.Memory.LimitInBytes)
	fmt.Println("Swap limit", recources.Memory.MemSwapLimitInBytes)
}
