package main

import (
	"fmt"
	"strings"
)

// Разворот слов в предложении

// Разработать программу, которая переворачивает порядок слов в строке.

// Пример: входная строка:

// «snow dog sun», выход: «sun dog snow».

// Считайте, что слова разделяются одиночным пробелом.
// Постарайтесь не использовать дополнительные срезы,
// а выполнять операцию «на месте».

func main() {
	fmt.Println(revertWords("1 2 3"))
}

func revertWords(words string) string {
	sepWords := strings.Split(words, " ")

	mid := len(sepWords) / 2
	for i := 0; i < mid; i++ {
		sepWords[i], sepWords[len(sepWords)-1-i] = sepWords[len(sepWords)-1-i], sepWords[i]
	}

	return strings.Join(sepWords, " ")
}
