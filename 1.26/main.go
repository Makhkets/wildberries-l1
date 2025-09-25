package main

import (
	"fmt"
	"strings"
)

// Уникальные символы в строке

// Разработать программу, которая проверяет, что все
// символы в строке встречаются один раз
// (т.е. строка состоит из уникальных символов).

// Вывод: true, если все символы уникальны,
// false, если есть повторения.
// Проверка должна быть регистронезависимой,
// т.е. символы в разных регистрах считать одинаковыми.

// Например: "abcd" -> true, "abCdefAaf" -> false
// (повторяются a/A), "aabcd" -> false.

// Подумайте, какой структурой данных
// удобно воспользоваться для проверки условия.

func unique(symbols string) bool {
	symbols = strings.ToLower(symbols)

	hashMap := make(map[rune]bool)

	for _, char := range []rune(symbols) {
		if hashMap[char] {
			return false
		}
		hashMap[char] = true
	}

	return true
}

func main() {
	symbols := "abCdefAaf"
	fmt.Println(unique(symbols))
}
