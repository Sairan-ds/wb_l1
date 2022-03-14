package main

import (
	"fmt"
	"time"
)

func main() {
	t := 3
	dur := time.Duration(t) * time.Second
	start := time.Now()
	sleep(dur)
	end := time.Since(start)
	fmt.Println("time of sleeping: ", end)
}
// функция для засыпания.
func sleep(t time.Duration) {
	<- time.After(t)
}