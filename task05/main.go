package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	timer := make(chan int)
	// Время работы программы
	t := 10

	// функция читает данные с канала ch
	go func(ch chan int) {
		for range ch {
			read := <-ch
			fmt.Println(read)
			t -= 1
		}
	}(ch)
	// функция в отдельном потоке включает таймер, по окончании которого отправляет сигнал в timer
	go func(timer chan int) {
		time.Sleep(time.Duration(t) * time.Second)
		timer <- 1
	}(timer)

	// Цикл в основном потоке который ждет сигнал от timer, при котором выйдет из цикла, если таймер не закончился, отправляются данные в канал ch
Loop:
	for {
		time.Sleep(time.Second)
		select {
		case <-timer:
			close(ch)
			break Loop
		default:
			ch <- t
		}
	}
}
