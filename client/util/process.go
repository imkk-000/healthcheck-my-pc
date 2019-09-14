package util

import (
	"healthcheck/client/model"

	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
)

func GetProcessByName(procs []*process.Process, name string) ([]*process.Process, error) {
	var filterProcs []*process.Process
	for _, proc := range procs {
		if procName, err := proc.Name(); err != nil {
			return nil, err
		} else if procName == name {
			filterProcs = append(filterProcs, proc)
		}
	}
	return filterProcs, error(nil)
}

func GetInfo(procs []*process.Process) model.HostInfo {
	bootTime, _ := host.BootTime()
	upTime, _ := host.Uptime()
	memInfoStat, _ := mem.VirtualMemory()
	var processesInfo []model.ProcessInfo
	for _, proc := range procs {
		var processInfo model.ProcessInfo
		processInfo.Name, _ = proc.Name()
		processInfo.CPUPercent, _ = proc.CPUPercent()
		processInfo.CPUAffinity, _ = proc.CPUAffinity()
		processInfo.Cmdline, _ = proc.Cmdline()
		processInfo.Exe, _ = proc.Exe()
		processInfo.MemoryInfoStat, _ = proc.MemoryInfo()
		processInfo.ConnStats, _ = proc.Connections()

		processesInfo = append(processesInfo, processInfo)
	}
	return model.HostInfo{
		BootTime:          bootTime,
		Uptime:            upTime,
		VirtualMemoryStat: *memInfoStat,
		ProcessesInfo:     processesInfo,
	}
}
