package process_test

import (
	. "healthcheck/util"
	"testing"

	"github.com/shirou/gopsutil/process"
	"github.com/stretchr/testify/assert"
)

func TestGetProcessByName(t *testing.T) {
	allProcs, err := process.Processes()
	assert.Nil(t, err)
	assert.NotEmpty(t, allProcs)

	filterProcsByName, err := GetProcessByName("go.exe")
	assert.Nil(t, err)
	assert.NotEmpty(t, filterProcsByName)

	assert.Greater(t, len(allProcs), len(filterProcsByName))
	for _, proc := range filterProcsByName {
		assert.Contains(t, allProcs, proc)
	}
}
