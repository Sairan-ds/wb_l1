package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	// делим строку по словам и создаем массив из этих слов
	tempArr := strings.Split(str, " ")
	arr := make([]string, 0)
	// добавляем по слову с конца массива в новый массив
	for i:= len(tempArr) - 1; i>=0; i -- {
		arr = append(arr, tempArr[i])
	}
	// преобразуем массив в строку
	str2 := strings.Join(arr, " ")
	fmt.Println(str, " - ", str2)
}