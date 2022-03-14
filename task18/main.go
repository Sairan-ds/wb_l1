package main

import (
	"fmt"
	"sync"
)
// структура счетчик
type incStruct struct {
	mx sync.Mutex 
	inc uint
	wg sync.WaitGroup
}

func newIncStruct() incStruct {
	return incStruct{}
}

func main() {
	// создаем нашу структуру
	inc := newIncStruct()
	// сколько раз будем выполнять работу
	times := 15
	// запуск конкурентных горутин которые будут писать конкурентно в счетчик
	for i:= 0; i < times; i ++ {
		inc.wg.Add(1)
		go makeSomeJob(i, &inc)
	}

	inc.wg.Wait()
	fmt.Printf("Работа была выполнена %v раз \n", inc.inc)
}


func makeSomeJob(i int,inc *incStruct)   {
	inc.mx.Lock()
	defer inc.mx.Unlock()
	fmt.Printf("Goroutine #%v make some job\n", i)
	inc.inc ++
	inc.wg.Done()
}