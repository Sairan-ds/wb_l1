package main

import (
	"fmt"
	"sync"
)

func main() {
	m := [...]int{2, 4, 6, 8, 10} 	// Массив из задания
	total := 0
	var wg sync.WaitGroup

	// Цикл по данному массиву с вызовом функции square в отдельной горутине
	for i := range m {
		wg.Add(1)
		go square(&total, m[i], &wg)
	}
	wg.Wait()
	fmt.Println(total)
}
// Функция добавляет квадрат принимаего числа в total (сумма квадратов)
func square(total *int, num int, wg *sync.WaitGroup) {
	*total += num*num
	wg.Done()
}
