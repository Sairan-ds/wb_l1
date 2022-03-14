package main

import "fmt"

func main() {
	// инизицализируем два канала и массив
	ch1 := make(chan int)
	ch2 := make(chan int)
	m := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

// считываем данные из массива и отправляем их в канал 1
	go func() {
		for _,i := range m {
			ch1 <- i
		}

	}()
// считываем данные из 1 канала и умножая на 2, отправляем во второй канал
	go func() {
		for range m{
			x := <-ch1
			ch2 <- x * 2
		}
		close(ch2)
	}()
 // Принимаем данные из 2 канала и выводим в stdout
	for {
		val , ok := <-ch2
		if ok {
			fmt.Println(val)
		} else {
			break
		}
	}
	close(ch1)
}
