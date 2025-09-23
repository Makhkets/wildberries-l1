package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на вход строку.

// Например: при вводе строки «главрыба» вывод должен быть «абырвалг».

// Учтите, что символы могут быть в Unicode (русские буквы, emoji и пр.),
// то есть просто iterating по байтам может не подойти — нужен срез рун ([]rune).

func main() {
	fmt.Println(revertWord("машина"))
}

// а ш а к
// 0 1 2 3

// к а ш а
// 0 1 2 3

func revertWord(word string) string {
	wordRune := []rune(word)

	mid := len(wordRune) / 2
	for i := 0; i < mid; i++ {
		wordRune[i], wordRune[len(wordRune)-1-i] = wordRune[len(wordRune)-1-i], wordRune[i]
	}

	return string(wordRune)
}
