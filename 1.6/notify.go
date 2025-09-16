package main

import (
	"fmt"
	"time"
)

// Остановка через канал уведомления
func main() {
	ch := make(chan int)
	notifyCh := make(chan bool)

	go func() {
		for {
			select {
			case <-notifyCh:
				return
			default:
				ch <- 1
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()

	time.Sleep(2 * time.Second)
	notifyCh <- true

	fmt.Println("остановили через канал уведомления")
}
