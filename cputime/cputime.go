// Package cputime measures CPU time of current thread.
package cputime

import "time"

// CurrentThread returns CPU time of current thread.
func CurrentThread() (time.Duration, error) {
	return currentThread()
}
