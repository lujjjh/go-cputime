# go-cputime

> Measure CPU time of current thread.

## Usage

```go
package main

import (
	"crypto/sha256"
	"log"
	"runtime"

	"github.com/lujjjh/go-cputime/cputime"
)

func main() {
	// Generally we should wire the goroutine to its current thread
	// so that we can measure accurate CPU time of the goroutine.
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	start, _ := cputime.CurrentThread()
	// Do something...
	for i := 0; i < 1e6; i++ {
		_ = sha256.Sum256([]byte("Hello, World!"))
	}
	end, _ := cputime.CurrentThread()

	// Calculate elapsed CPU time.
	log.Println("CPU time:", end-start)
}
```
