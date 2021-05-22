package main

import (
	"fmt"
	"os"
	"syscall"
)

// NOTE: Self implementation https://github.com/golang/go/blob/go1.16.4/src/os/exec_plan9.go#L142
func getExitStatus(state *os.ProcessState) int {
	if state == nil {
		return -1
	}

	status, ok := state.Sys().(syscall.WaitStatus)
	if !ok {
		fmt.Fprintf(stderr, "%s\n", "failed wait status assertion")
		return -1
	}

	return status.ExitStatus()
}
