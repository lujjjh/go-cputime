package cputime

// #include <time.h>
//
// typedef struct timespec timespec;
//
import "C"
import (
	"time"
)

func CurrentThread() (time.Duration, error) {
	var ts C.timespec
	_, err := C.clock_gettime(C.CLOCK_THREAD_CPUTIME_ID, &ts)
	if err != nil {
		return 0, err
	}
	// TODO: overflow?
	return time.Duration(int64(ts.tv_sec)*1e9 + int64(ts.tv_nsec)), nil
}
