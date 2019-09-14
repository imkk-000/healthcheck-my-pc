package model

import (
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

type HostInfo struct {
	BootTime          uint64                `json:"bootTime"`
	Uptime            uint64                `json:"uptime"`
	VirtualMemoryStat mem.VirtualMemoryStat `json:"virtualMemoryStat"`
	ProcessesInfo     []ProcessInfo         `json:"processesInfo"`
}

type ProcessInfo struct {
	Pid            int32                   `json:"pid"`
	Name           string                  `json:"name"`
	CPUPercent     float64                 `json:"cpuPercent"`
	CPUAffinity    []int32                 `json:"cpuAffinity"`
	Cmdline        string                  `json:"cmdline"`
	Exe            string                  `json:"exe"`
	ConnStats      []net.ConnectionStat    `json:"connStats"`
	MemoryInfoStat *process.MemoryInfoStat `json:"memoryInfoStat"`
}
