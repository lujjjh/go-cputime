package cputime

import (
	"runtime"
	"testing"
	"time"
)

func abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

func TestCurrentThread(t *testing.T) {
	done := make(chan struct{})
	go func() {
		defer close(done)
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()

		const (
			expected = 50 * time.Millisecond
			eps      = 25 * time.Millisecond
		)

		start, _ := CurrentThread()
		for s := time.Now(); time.Since(s) < expected; {
		}
		time.Sleep(expected)
		end, _ := CurrentThread()
		cpuTime := end - start
		if abs(int64(cpuTime-expected)) > int64(eps) {
			t.Errorf("expected %s, got %s", expected, cpuTime)
		}
	}()
	<-done
}
