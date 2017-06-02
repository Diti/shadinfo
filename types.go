package main

import (
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

type ShadInfo struct {
	Cpu  []ShadInfoCpu
	Disk *disk.UsageStat
	Host *host.InfoStat
	Mem  *mem.VirtualMemoryStat
	Net  []net.ConnectionStat
	Proc ShadInfoProc
}

type ShadInfoCpu struct {
	ModelName string
	Cores     int32
	Mhz       float64
	Percent   float64
}

type ShadInfoDisk disk.UsageStat
type ShadInfoHost host.InfoStat
type ShadInfoMem mem.VirtualMemoryStat
type ShadInfoNet net.ConnectionStat

type ShadInfoProc struct {
	command string
	memory  *process.MemoryInfoStat
	pid     string
}
