package main

import (
	"fmt"
	"time"
)

//////////////////////////////////
// второй вариант, с time.After //
//////////////////////////////////

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала – читать эти значения. По истечении N секунд программа
// должна завершаться.

// Подсказка: используйте time.After или таймер для ограничения времени работы.

func main() {
	ch := make(chan int)
	timer := time.After(3 * time.Second)
	done := make(chan bool)

	// writer
	go func() {
		defer func() { done <- true }()
		for {
			select {
			case <-timer:
				fmt.Println("завершаем writer")
				return
			default:
				ch <- 1
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// reader
	go func() {
		defer close(ch)
		for v := range ch {
			fmt.Println(v)
		}
	}()

	<-done
}
