package main

import "fmt"

func main() {
	str := "Главрыба"
	// создаем массив байтов нашей строки 
	tempAr := []rune(str)
	array := make([]rune, 0)
	// начинаем с конца читать наш массив в обратном порядке и добавлять буквы в итоговый массив
	for i :=  len(tempAr)-1; i >= 0; i -- {
		array = append(array, tempAr[i])
	}
	// преобразуем массив байтов в строку
	fmt.Println(string(array))
}