package main

import (
	"fmt"
	"time"
)

// Конвейер чисел

// Разработать конвейер чисел. Даны два канала:

// в первый пишутся числа x из массива,
// во второй – результат операции x*2. После этого данные из второго канала должны выводиться в stdout.

// То есть, организуйте конвейер из двух этапов с горутинами:
// генерация чисел и их обработка. Убедитесь, что чтение из второго канала корректно завершается.

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go gen(ch1)
	go work(ch1, ch2)

	go func() {
		for v := range ch2 {
			fmt.Println(v)
		}
	}()

	time.Sleep(1 * time.Second)
}

func gen(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func work(ch, ch2 chan int) {
	for v := range ch {
		ch2 <- v * 2
	}
	close(ch2)
}
