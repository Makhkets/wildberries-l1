package main

import "fmt"

var num int64 = 5

func setBit(n int64, pos int, value int) int64 {
	if value == 1 {
		return n | (1 << pos)
	}
	return n &^ (1 << pos)
}

func main() {
	fmt.Printf("Число: %d (двоичное: %b)\n", num, num)

	result1 := setBit(num, 0, 0)
	fmt.Printf("Устанавливаем 0-й бит в 0: %d (двоичное: %b)\n", result1, result1)

	result2 := setBit(num, 1, 1)
	fmt.Printf("Устанавливаем 1-й бит в 1: %d (двоичное: %b)\n", result2, result2)

	result3 := setBit(num, 2, 1)
	fmt.Printf("Устанавливаем 2-й бит в 1: %d (двоичное: %b)\n", result3, result3)
}
