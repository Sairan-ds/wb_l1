package main

import (
	"fmt"
	"sync"
)

func main() {
	m := [...]int{2, 4, 6, 8, 10} 	// Массив из задания

	var wg sync.WaitGroup

	// Цикл по данному массиву с вызовом функции square в отдельной горутине
	for i := range m {
		wg.Add(1)
		go square( m[i], &wg)
		wg.Wait()
	}
	


}
// Функция принимает на вход ссылку на WG и значение для которого отображает квадрат в stdout
func square(num int, wg *sync.WaitGroup) {
	fmt.Println(num*num)
	wg.Done()
}
