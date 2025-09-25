package main

import (
	"fmt"
	"math/big"
)

// Большие числа и операции

// Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a, b, значения которых > 2^20 (больше 1 миллион).

//  Комментарий: в Go тип int справится с такими числами, но обратите
//  внимание на возможное переполнение для ещё больших значений.
//  Для очень больших чисел можно использовать math/big.

func main() {
	a, _ := big.NewInt(0).SetString("10000000000000000000000", 10)
	b, _ := big.NewInt(0).SetString("5000000000000000000000", 10)

	fmt.Println(add(a, b))
	fmt.Println(subtract(a, b))
	fmt.Println(multiply(a, b))
	fmt.Println(divide(a, b))
}

func add(a, b *big.Int) *big.Int {
	return big.NewInt(0).Add(a, b)
}

func subtract(a, b *big.Int) *big.Int {
	return big.NewInt(0).Sub(a, b)
}

func multiply(a, b *big.Int) *big.Int {
	return big.NewInt(0).Mul(a, b)
}

func divide(a, b *big.Int) *big.Int {
	return big.NewInt(0).Div(a, b)
}
