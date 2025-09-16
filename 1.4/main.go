package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).

// Выберите и обоснуйте способ завершения работы всех горутин-воркеров при получении сигнала
// прерывания.

// Подсказка: можно использовать контекст (context.Context) или канал для оповещения о завершении.

func main() {
	context, cancel := context.WithCancel(context.Background())
	done := make(chan bool)

	go temp(context, done)

	//time.Sleep(3 * time.Second)
	//cancel()

	quit := make(chan os.Signal, 1)

	// sigint - прерывание, sigterm - завершение
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	cancel()
	<-done

	fmt.Println("завершение")
}

func temp(context context.Context, done chan bool) {
	defer func() { done <- true }()

	for {
		select {
		case <-context.Done():
			fmt.Println("завершение горутины")
			return
		default:
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
