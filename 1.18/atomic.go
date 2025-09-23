package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var counter atomic.Int32

func main() {
	for range 10 {
		go func() {
			counter.Add(1)
		}()
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Print(counter.Load())
}
