package main

import (
	// "context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// Первый способ связан с отправкой любого сигнала в select, в котором будет выход из цикла и завершение горутины
	// Пример реализации
	fmt.Println("1 способ")
	ch := make(chan bool)
	closedChan := make(chan bool)
	go func() {
		for {
			select {
			case <-closedChan:
				fmt.Println("Closed")
				close(ch)
				return
			default:
				fmt.Println("Not closed")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	closedChan <- true
	// Либо мы можем просто закрыть канал и горутина также завершится
	//close(closedChan)
	// Данный канал позволит дождаться завершения горутины
	<-ch

	fmt.Println("2 способ")
	var wg1 sync.WaitGroup
	ch2 := make(chan bool)
	go func() {
		defer wg1.Done()
		for i := range ch2 {
			fmt.Println("Some work", i)
		}
	}()

	ch2 <- true
	time.Sleep(3 * time.Second)
	close(ch2)
	wg1.Wait()

	// 3 способ, аналог первого через ctx
	// fmt.Println("3 способ")
	// forever := make(chan struct{})
	// ctx, cancel := context.WithCancel(context.Background())

	// go func(ctx context.Context) {
	// 	for {
	// 		select {
	// 		case <-ctx.Done(): // if cancel() execute
	// 			forever <- struct{}{}
	// 			return
	// 		default:
	// 			fmt.Println("for loop")
	// 		}

	// 		time.Sleep(500 * time.Millisecond)
	// 	}
	// }(ctx)

	// go func() {
	// 	time.Sleep(3 * time.Second)
	// 	cancel()
	// }()

	// <-forever
	// fmt.Println("finish")

}
