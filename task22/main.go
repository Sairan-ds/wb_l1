package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	a.SetString("34000000000000000000", 10)

	b := new(big.Int)
	b.SetString("1000000000000000000", 10)

	c := big.NewInt(0)

	c.Div(a, b)
	fmt.Printf("Деление  %22v и %v = %v \n", a, b, c)

	c.Mul(b, a)
	fmt.Printf("Умножение %21v и %v = %v \n", a, b, c)

	c.Add(b, a)
	fmt.Printf("Сумма %25v и %v = %v \n", a, b, c)

	c.Sub(b, a)
	fmt.Printf("Разность %22v и %v = %v \n", a, b, c)
}
