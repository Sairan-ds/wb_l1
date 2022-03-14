package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"
	fmt.Println(checkUnice(str1))
	fmt.Println(checkUnice(str2))
	fmt.Println(checkUnice(str3))
}
// Так как функция не регистрозависимая, то переводим строку в нижний регистр. Затем проверяем уникальность. Если все ок отправляем True. 
func checkUnice(str string) (string ,bool) {
	str = strings.ToLower(str)
	tempMap := make(map[rune]int)
	for _, i := range str {
		tempMap[i] += 1
	}
	for _, v := range tempMap {
		if v > 1 {
			return str + " -", false
		}
	}
	return str + " -", true
}