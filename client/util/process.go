package process

import (
	"github.com/shirou/gopsutil/process"
)

func GetProcessByName(name string) ([]*process.Process, error) {
	procs, err := process.Processes()
	if err != nil {
		return nil, err
	}
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
