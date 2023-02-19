package gojue

import "golang.org/x/sys/unix"

type goJueConfig struct {
	Kernelmajor byte
	Kernelminor byte
	Kernelpatch byte
	btfEnable   bool
	enableCORE  bool
}

type eBPFConfig struct {

	// PerfRingBufferSize - Manager-level default value for the perf ring buffers. Defaults to the size of 1 page
	perCPUBuffer int

	// SymFile - Kernel symbol file. If not provided, the default `/proc/kallsyms` will be used.
	SymFile int

	// RLimit - The maps & programs provided to the manager might exceed the maximum allowed memory lock.
	// (RLIMIT_MEMLOCK) If a limit is provided here it will be applied when the manager is initialized.
	RLimit *unix.Rlimit
}
