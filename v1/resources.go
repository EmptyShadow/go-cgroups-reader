package cgroupsreader

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/EmptyShadow/go-cgroups-reader/internal"
)

const (
	sysDir     = "sys"
	fsDir      = "fs"
	cgroupsDir = "cgroup"
	cpuDir     = "cpu"
	memoryDir  = "memory"

	// https://www.kernel.org/doc/Documentation/scheduler/sched-bwc.txt.
	cpuCFSPeriodUSFile = "cpu.cfs_period_us" // the total available run-time within a period (in microseconds).
	cpuCFSQuotaUSFile  = "cpu.cfs_quota_us"  // the length of a period (in microseconds)

	// https://www.kernel.org/doc/Documentation/cgroup-v1/memory.txt.
	memoryLimitInBytesFile        = "memory.limit_in_bytes"       // limit of memory usage.
	memoryMemSwapLimitInBytesFile = "memory.memsw.limit_in_bytes" // limit of memory+Swap usage.
)

var (
	cgroupDir = filepath.Join(sysDir, fsDir, cgroupsDir)

	cpuCFSPeriodUSFilePath = filepath.Join(cpuDir, cpuCFSPeriodUSFile)
	cpuCFSQuotaUSFilePath  = filepath.Join(cpuDir, cpuCFSQuotaUSFile)

	memoryLimitInBytesFilePath        = filepath.Join(memoryDir, memoryLimitInBytesFile)
	memoryMemSwapLimitInBytesFilePath = filepath.Join(memoryDir, memoryMemSwapLimitInBytesFile)
)

type Resources struct {
	CPU    CPUResources
	Memory MemoryResources
}

type CPUResources struct {
	PeriodUS time.Duration
	QuotaUS  time.Duration
}

type MemoryResources struct {
	LimitInBytes        uint64
	MemSwapLimitInBytes uint64
}

func ReadResources(root, proc string) (resources Resources, enabled bool, err error) {
	procDir := filepath.Join(root, cgroupDir, proc)

	if !internal.IsExistsFile(procDir) {
		return resources, false, nil
	}

	resources.CPU, err = readCPUResources(procDir)
	if err != nil {
		return resources, false, fmt.Errorf("load cpu resources: %w", err)
	}

	resources.Memory, err = readMemoryResources(procDir)
	if err != nil {
		return resources, false, fmt.Errorf("load memory resources: %w", err)
	}

	return resources, true, nil
}

func ReadCPUResources(root, proc string) (resources CPUResources, enabled bool, err error) {
	procDir := filepath.Join(root, cgroupDir, proc)

	if !internal.IsExistsFile(procDir) {
		return resources, false, nil
	}

	resources, err = readCPUResources(procDir)

	return resources, true, err
}

func readCPUResources(root string) (resources CPUResources, err error) {
	cpuPeriodUS, err := internal.ReadInt64FromFile(filepath.Join(root, cpuCFSPeriodUSFilePath))
	if err != nil {
		return resources, fmt.Errorf("read %s: %w", cpuCFSPeriodUSFilePath, err)
	}

	cpuQuotaUS, err := internal.ReadInt64FromFile(filepath.Join(root, cpuCFSQuotaUSFilePath))
	if err != nil {
		return resources, fmt.Errorf("read %s: %w", cpuCFSQuotaUSFilePath, err)
	}

	resources.PeriodUS = time.Microsecond * time.Duration(cpuPeriodUS)
	resources.QuotaUS = time.Microsecond * time.Duration(cpuQuotaUS)

	return
}

func ReadMemoryResources(root, proc string) (resources MemoryResources, enabled bool, err error) {
	procDir := filepath.Join(root, cgroupDir, proc)

	if !internal.IsExistsFile(procDir) {
		return resources, false, nil
	}

	resources, err = readMemoryResources(procDir)

	return resources, true, err
}

func readMemoryResources(root string) (resources MemoryResources, err error) {
	memoryLimitInBytes, err := internal.ReadUInt64FromFile(filepath.Join(root, memoryLimitInBytesFilePath))
	if err != nil {
		return resources, fmt.Errorf("read %s: %w", cpuCFSQuotaUSFilePath, err)
	}

	memoryMemSwapLimitInBytes, err := internal.ReadUInt64FromFile(filepath.Join(root, memoryMemSwapLimitInBytesFilePath))
	if err != nil {
		return resources, fmt.Errorf("read %s: %w", cpuCFSQuotaUSFilePath, err)
	}

	resources.LimitInBytes = memoryLimitInBytes
	resources.MemSwapLimitInBytes = memoryMemSwapLimitInBytes

	return
}
