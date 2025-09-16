package main

import (
	"context"
	"fmt"
	"time"
)

// остановка через context
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("завершаем горутину через context")
				return
			}
		}
	}()

	cancel()
	time.Sleep(1 * time.Second)
}
