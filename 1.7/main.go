package main

import (
	"sync"
	"time"
)

// # Конкурентная запись в map	

// Реализовать безопасную для конкуренции запись данных в структуру map.

// Подсказка: необходимость использования синхронизации
// (например, sync.Mutex или встроенная concurrent-map).

// Проверьте работу кода на гонки (util go run -race).

var data struct {
	sync.Mutex
	m map[string]interface{}
} = struct {
	sync.Mutex
	m map[string]interface{}
}{m: make(map[string]interface{})}

func main() {
	done := make(chan bool)
	for i := 0; i < 3; i++ {
		go func() {
			for {
				data.Lock()
				data.m["time"] = time.Now()
				time.Sleep(100 * time.Millisecond)
				data.Unlock()
			}
		}()
	}
	<-done
}
