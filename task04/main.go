package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var N int
	fmt.Print("Введите количество воркеров: ")
	fmt.Fscan(os.Stdin, &N)

	ch := make(chan int)
	defer close(ch)

	work := func(id int, ch <- chan int) {
		for {
			value := <- ch
			fmt.Printf("Worker #%v receive number %v \n", id, value)
		}
	}

	for i := 0; i < N; i++ {
		go work(i, ch)
	}

	for i:= 0; i >= 0; i ++{
		time.Sleep(2*time.Second)
		ch <- i
	}

}
