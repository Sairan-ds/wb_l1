package main

import "fmt"

func main() {
	//  cgb
	m := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	mm := make(map[int][]float64)
	for _,i := range m {
		key := int(i) / 10 * 10
		mm[key] = append(mm[key], i)
	}

	fmt.Println(mm)
}
