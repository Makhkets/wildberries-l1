package main

import (
	"fmt"
)

// Бинарный поиск

// Реализовать алгоритм бинарного поиска встроенными методами языка.
// Функция должна принимать отсортированный слайс и искомый элемент,
// возвращать индекс элемента или -1, если элемент не найден.

// Подсказка: можно реализовать рекурсивно или итеративно, используя цикл for.

func main() {
	arr := []int{}
	for i := 0; i < 100+1; i++ {
		arr = append(arr, i)
	}

	fmt.Println("result", binarySearch(arr, 11))
}

func binarySearch(arr []int, target int) int {
	if target > arr[len(arr)-1] || target < arr[0] {
		return -1
	}

	center := arr[len(arr)/2]
	fmt.Println("center", center)

	if target == center {
		return target
	} else {
		if center > target {
			return binarySearch(arr[:len(arr)/2], target)
		} else {
			return binarySearch(arr[len(arr)/2:], target)
		}
	}
}
