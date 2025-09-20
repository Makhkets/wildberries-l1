package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов) — т.е.
// вывести элементы, присутствующие и в первом, и во втором.

// Пример:
// A = {1,2,3}
// B = {2,3,4}
// Пересечение = {2,3}

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	hMap := map[int]bool{}
	result := []int{}

	for _, v := range A {
		hMap[v] = true
	}

	for _, v := range B {
		if hMap[v] {
			result = append(result, v)
		}
	}

	fmt.Println("результат", result)
}
