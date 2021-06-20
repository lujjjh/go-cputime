package cputime

import (
	"syscall"
	"time"
	"unsafe"
)

const clockThreadCputimeId = 3

func currentThread() (time.Duration, error) {
	var ts syscall.Timespec
	_, _, err := syscall.RawSyscall(syscall.SYS_CLOCK_GETTIME, uintptr(clockThreadCputimeId), uintptr(unsafe.Pointer(&ts)), 0)
	if err != 0 {
		return 0, err
	}
	return time.Duration(ts.Nano()), nil
}
