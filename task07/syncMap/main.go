package main

import (
	"fmt"
	"sync"
)
// sync.Map решает совершенно конкретную проблему cache contention в стандартной библиотеке для таких случаев,
// когда ключи в map стабильны (не обновляются часто) и происходит намного больше чтений, чем записей.
func main() {
	var counters sync.Map
// Сохранение данных
	counters.Store("test", 111)
// Чтение данных
	val, ok := counters.Load("test")

	fmt.Println(val, ok)
}
