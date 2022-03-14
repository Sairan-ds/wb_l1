package main

import (
	"fmt"
	"sync"
)

// Создаем новую структуру данных и встраиваем в нее mutex
type Counters struct {
	mx sync.Mutex
	m map[string]int
}
// Конструктор для инициализации нашей структуры, мбютекс инициализировать не надо, его нулевое значение - разлоченный мьютекс
func NewCounters() *Counters {
    return &Counters{
        m: make(map[string]int),
    }
}
// Чтение данных из мапы
func (c *Counters) Load(key string) (int, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	val, ok := c.m[key]
	return val, ok
}
// Сохранение в мапу
func (c *Counters) Store(key string, value int) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key] = value
}

func main() {
	c := NewCounters()
	c.Store("test", 111)
	v,ok := c.Load("test")
	fmt.Println(v, ok)
}