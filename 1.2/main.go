package main

import (
	"fmt"
	"sync"
)

var m = []int{2, 4, 6, 8, 10}

func main() {
	var wg sync.WaitGroup
	for _, v := range m {
		wg.Go(func() {
			fmt.Println(v * v)
		})
	}
	wg.Wait()
}
