package main

import (
	"fmt"
	"sync"
)
// Чтение, как мы помним, из map — безопасно, но сейчас мы будем лочиться на каждой операции чтения, теряя без особой выгоды время на ожидание разблокировки.
// В стандартной библиотеке для решения этой ситуации есть тип sync.RWMutex. Помимо Lock()/Unlock(), 
// у RWMutex есть отдельные аналогичные методы только для чтения — RLock()/RUnlock()
// Создаем новую структуру данных и встраиваем в нее mutex
type Counters struct {
	mx sync.RWMutex
	m  map[string]int
}
// Конструктор для инициализации нашей структуры, мбютекс инициализировать не надо, его нулевое значение - разлоченный мьютекс
func NewCounters() *Counters {
    return &Counters{
        m: make(map[string]int),
    }
}
// Чтение данных из мапы
func (c *Counters) Load(key string) (int, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
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
	v, ok := c.Load("test")
	fmt.Println(v, ok)
}
