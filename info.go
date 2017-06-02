package main

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

func PopulateInfo() (info *ShadInfo) {
	// Get system information with `gopsutil`
	cpuInfo, err := cpuInfo(false)
	if err != nil {
		logPanic(err)
		cpuInfo = nil
	}
	diskInfo, err := diskInfo()
	if err != nil {
		logPanic(err)
		diskInfo = nil
	}
	hostInfo, err := hostInfo()
	if err != nil {
		logPanic(err)
		hostInfo = nil
	}
	memInfo, err := memInfo()
	if err != nil {
		memInfo = nil
		logPanic(err)
	}

	// Populate struct with info
	info = &ShadInfo{
		Cpu:  cpuInfo,
		Disk: diskInfo,
		Host: hostInfo,
		Proc: ShadInfoProc{},
		Mem:  memInfo,
		Net:  nil,
	}
	return
}

// CPU information.
// If `combined` is set to false, this function returns each CPU independently.
func cpuInfo(combined bool) (cpuInfo []ShadInfoCpu, err error) {
	gopsInfo, err := cpu.Info()
	if err != nil {
		return
	}
	cpuPercent, err := cpu.Percent(0, !combined)
	if err != nil {
		return
	}
	cpuCount, err := cpu.Counts(false)
	if err != nil {
		return
	}

	cpuInfo = make([]ShadInfoCpu, cpuCount)
	for i := 0; i < cpuCount; i++ {
		cpuInfo[i] = ShadInfoCpu{
			ModelName: gopsInfo[0].ModelName,
			Cores:     gopsInfo[0].Cores,
			Mhz:       gopsInfo[0].Mhz,
			Percent:   cpuPercent[i],
		}
	}
	return
}

func diskInfo() (*disk.UsageStat, error) {
	info, err := disk.Usage(".")
	if err != nil {
		return &disk.UsageStat{}, err
	}
	return info, nil
}

func hostInfo() (*host.InfoStat, error) {
	info, err := host.Info()
	if err != nil {
		return &host.InfoStat{}, err
	}
	return info, nil
}

func memInfo() (*mem.VirtualMemoryStat, error) {
	info, err := mem.VirtualMemory()
	if err != nil {
		return &mem.VirtualMemoryStat{}, err
	}
	return info, nil
}

func netInfo() ([]net.ConnectionStat, error) {
	info, err := net.Connections("inet")
	if err != nil {
		return []net.ConnectionStat{}, err
	}
	return info, nil
}

func processInfo() ([]ShadInfoProc, error) {
	info := []ShadInfoProc{}

	pids, err := process.Pids()
	if err != nil {
		return []ShadInfoProc{}, err
	}

	for _, pid := range pids {
		pr := ShadInfoProc{}
		process, err := process.NewProcess(pid)
		if err != nil {
			return []ShadInfoProc{}, err
		}

		// No error checking because it somehow always returns error “exit status 1”
		cmd, _ := process.Cmdline()
		mem, _ := process.MemoryInfo()

		pr.command = cmd
		pr.memory = mem
		pr.pid = fmt.Sprint(pid)

		info = append(info, pr)
	}
	return info, nil
}

func tempInfo() ([]host.TemperatureStat, error) {
	info, err := host.SensorsTemperatures()
	if err != nil {
		return []host.TemperatureStat{}, err
	}
	return info, nil
}
