package main

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func cpuInfo() ([]cpu.InfoStat, error) {
	info, err := cpu.Info()
	if err != nil {
		return []cpu.InfoStat{}, err
	}
	return info, nil
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
	info, err := net.Connections("all")
	if err != nil {
		return []net.ConnectionStat{}, err
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
