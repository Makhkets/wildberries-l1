package main

import (
	"context"
	"fmt"
	"time"
)

//////////////////////////////////////////////////
// первый вариант, без time.After				//
// вариант с time.After см. 1.5/1.5_2/main.go   //
//////////////////////////////////////////////////

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала – читать эти значения. По истечении N секунд программа
// должна завершаться.

// Подсказка: используйте time.After или таймер для ограничения времени работы.

func main() {
	ch := make(chan string)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	go writer(ch, ctx)
	go reader(ch, ctx)

	<-ctx.Done()

	fmt.Println("\nзавершаем")
}

func writer(ch chan string, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("writer: <-ctx.Done()")
			return
		default:
			ch <- "."
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func reader(ch chan string, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("reader: <-ctx.Done()")
			return
		default:
			fmt.Print(<-ch)
		}
	}
}
