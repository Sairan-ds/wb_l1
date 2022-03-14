package main

import (
	"fmt"
	"math"
)

// точка
type point struct {
	x,y float64
}

// конструктор точки
func newPoint(a,b float64) *point {
	return &point{
		x: a,
		y: b,
	}
}
// функция рассчета расстояния между точками
func distance(a,b *point) float64 {
	return math.Round(math.Sqrt(math.Pow(a.x - b.x, 2) + math.Pow(a.y - b.y, 2))*100)/100
}

func main() {
	a := newPoint(1,2)
	b := newPoint(3,6)

	c := distance(a,b)

	fmt.Println(c)
}