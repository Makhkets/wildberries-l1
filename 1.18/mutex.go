package main

import (
	"fmt"
	"sync"
	"time"
)

// Конкурентный счетчик

// Реализовать структуру-счётчик,
// которая будет инкрементироваться в конкурентной среде
// (т.е. из нескольких горутин).
// По завершению программы структура должна выводить итоговое значение счётчика.

// Подсказка: вам понадобится механизм синхронизации,
// например, sync.Mutex или sync/Atomic для безопасного инкремента.

var Counter = struct {
	Inc int
	mu  sync.Mutex
}{}

func main() {
	for i := 0; i < 3; i++ {
		go func() {
			Counter.mu.Lock()
			Counter.Inc++
			Counter.mu.Unlock()
		}()
	}

	time.Sleep(200 * time.Millisecond)
	fmt.Println(Counter.Inc)
}
