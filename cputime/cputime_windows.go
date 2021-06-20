package cputime

import (
	"syscall"
	"time"
	"unsafe"
)

var (
	kernel32         = syscall.MustLoadDLL("kernel32.dll")
	getCurrentThread = kernel32.MustFindProc("GetCurrentThread")
	getThreadTimes   = kernel32.MustFindProc("GetThreadTimes")
)

func currentThread() (time.Duration, error) {
	handle, _, err := getCurrentThread.Call()
	if handle == 0 {
		return 0, err
	}
	var (
		creationTime syscall.Filetime
		exitTime     syscall.Filetime
		kernelTime   syscall.Filetime
		userTime     syscall.Filetime
	)
	r, _, err := getThreadTimes.Call(
		handle,
		uintptr(unsafe.Pointer(&creationTime)),
		uintptr(unsafe.Pointer(&exitTime)),
		uintptr(unsafe.Pointer(&kernelTime)),
		uintptr(unsafe.Pointer(&userTime)))
	if r == 0 {
		return 0, err
	}
	return time.Duration(kernelTime.Nanoseconds() + userTime.Nanoseconds()), nil
}
