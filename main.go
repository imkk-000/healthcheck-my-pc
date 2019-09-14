package main

import (
	"fmt"

	"github.com/shirou/gopsutil/process"
)

func main() {
	procs, err := process.Processes()
	if err != nil {
		panic(err)
	}
	for i, proc := range procs {
		name, _ := proc.Name()
		conn, _ := proc.Connections()
		fmt.Println(i, name, conn)
	}
}
