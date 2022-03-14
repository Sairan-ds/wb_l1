package main

import (
	"fmt"
)

var justString string
// Под капотом Го фактически кодирует строку как массив байтов. Но не различает строки ASCII  и Unicode. 
// Так как некоторые символы занимают 2 или 3 байта(к примеру китайские символы). 
// Пример сделан на генерации 1024 символов буквы Д, которая занимает два байта

// Генерация строки 
func createHugeString(size int) (s string) {
	for i := 0; i < size; i++ {
		s+= "Д"
	}
	return 
}
// Вывод первых 100 символов
func someFunc() {
	v := createHugeString(1 << 10)
	a := []rune(v)
	// Здесь мы сразу видим, что количество символов различается в байтах и рунах. Поэтому, чтобы у нас было правильно отображение, 
	// нужно переводить строку в руны. 
	fmt.Println(len(v))
	fmt.Println(len(a))

	justString = string(a[:100])

	fmt.Println(justString)

}

func main() {
	someFunc()


}
