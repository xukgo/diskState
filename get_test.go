package main

import (
	"fmt"
	"testing"
)

// example
func TestGet(t *testing.T) {
	state := DiskUsage("/")
	fmt.Printf("All=%dM, Free=%dM, Available=%dM, Used=%dM, Usage=%d%%",
		state.All/MB, state.Free/MB, state.Available/MB, state.Used/MB, 100*state.Used/state.All)
}
