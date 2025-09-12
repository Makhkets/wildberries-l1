package main

import "fmt"

type Human struct {
	Name string
	Age  uint
}

type Action struct {
	Human
}

func (h *Human) SayHello() string {
	return "my name is " + h.Name
}

func main() {
	s := Action{Human{"Makhkets", 25}}
	fmt.Print(s.SayHello())
}
