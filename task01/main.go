package main

import "fmt"
// Родительская структура Human
type Human struct {
	Name string
	Age  uint16
}
// Структура где будет реализовоно встраивание методов от родительской структуры Human (аналог наследования).
type Action struct {
	Human
	Breath string
	Run    string
}
// Метод родительской структуры Human
func (h Human) Say() {
	fmt.Printf("Hi, I am %s, and I am %v years old \n", h.Name, h.Age)
}

func main() {
	// Создадим Action
	a := Action{
		Human: Human{
			Name: "John",
			Age:  28,
		},
		Breath: "breath",
		Run:    "fast",
	}
	// Вызываем метод родительской структуры Human
	a.Say()
}
