package main

import "fmt"

// # Собственное множество строк

// Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree").
// Создать для неё собственное множество.

// Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	hmap := map[string]bool{}
	res := []string{}

	for i := 0; i < len(words); i++ {
		if hmap[words[i]] {
			continue
		} else {
			hmap[words[i]] = true
			res = append(res, words[i])
		}
	}

	fmt.Println(res)
}
