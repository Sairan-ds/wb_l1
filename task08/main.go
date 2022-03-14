package main

import "fmt"

func main() {
	// Задаем число в котором будем устанавливать  i-ый бит в 1 или 0
	var val int64 = 63
	fmt.Printf("%b\n", val)
	// Если будем менять i-ый бит на 1, то ставим true, в другом случае false
	valBool := false
	// Выбор бита который будем менять
	valByte := 1
	if valBool {
		val |= 1 << valByte
	} else {
		val &^= 1 << valByte
	}

	fmt.Printf("%b\n", val)
}