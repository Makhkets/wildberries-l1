package main

import "fmt"

// Реализовать алгоритм быстрой сортировки массива встроенными средствами языка.
// Можно использовать рекурсию.

// Подсказка: напишите функцию quickSort([]int) []int которая сортирует срез целых чисел.
// Для выбора опорного элемента можно взять середину или первый элемент.

// P.S Я выбрал метод сортировки пузырьком
// https://www.youtube.com/watch?v=PF7AqefS4MU (2:10 - 4:00)

func main() {
	fmt.Println("Sorted array:", bubbleSort([]int{5, 3, 8, 6, 2}))
}

func bubbleSort(arr []int) []int {
	arrLen := len(arr) - 1
	for arrLen != 0 {
		maxIndex := 0
		for i := 0; i < arrLen; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
			maxIndex = i
		}

		arrLen = maxIndex
	}

	return arr
}
