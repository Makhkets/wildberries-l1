package wildberries_l1

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер %d завершается\n", id)
			return
		default:
			// Имитация работы
			fmt.Printf("Воркер %d работает\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// Запускаем несколько воркеров
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, &wg, i)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	fmt.Println("Получен сигнал завершения")

	// Отменяем контекст, что сигнализирует всем горутинам о завершении
	cancel()

	// Ждем завершения всех горутин
	wg.Wait()
	fmt.Println("Все горутины завершены")
}
