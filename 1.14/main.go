package main

import "fmt"

// # Определение типа переменной в runtime

// Разработать программу, которая в runtime способна определить тип переменной,
// переданной в неё (на вход подаётся interface{}).
// Типы, которые нужно распознавать: int, string, bool, chan (канал).

// Подсказка: оператор типа switch v.(type) поможет в решении.

func main() {
	str := "test"
	integer := 1
	boolean := true
	channel := make(chan interface{})

	getType(str)
	getType(integer)
	getType(boolean)
	getType(channel)
}

func getType(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	case chan interface{}:
		fmt.Println("chan")
	default:
		fmt.Println("unknown type")
	}
}
