package util_test

import (
	"healthcheck/client/model"
	. "healthcheck/client/util"
	"testing"

	"github.com/shirou/gopsutil/process"
	"github.com/stretchr/testify/assert"
)

func TestGetProcessesInfo(t *testing.T) {
	procs, _ := process.Processes()
	filterProcsByName, _ := GetProcessByName(procs, "go")
	hostInfo := GetInfo(filterProcsByName)
	assert.NotNil(t, hostInfo)
	assert.NotEmpty(t, hostInfo)
	assert.IsType(t, model.HostInfo{}, hostInfo)
}

func TestGetProcessByName(t *testing.T) {
	procs, err := process.Processes()
	assert.Nil(t, err)
	assert.NotEmpty(t, procs)

	filterProcsByName, err := GetProcessByName(procs, "go")
	assert.Nil(t, err)
	assert.NotEmpty(t, filterProcsByName)

	assert.Greater(t, len(procs), len(filterProcsByName))
	for _, proc := range filterProcsByName {
		assert.Contains(t, procs, proc)
	}
}
