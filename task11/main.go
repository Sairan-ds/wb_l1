package main

import "fmt"

func main() {
	// Реализация пересечения двух множест
	input1 := []int{1, 2, 3, 4, 5}
	input2 := []int{4, 5, 6}

	interceptionSet(input1, input2)

	// Реализация пересечения двух слайсов, где могут повторяться значения
	// tempInt := make([]int, 0)
	// 	tempMap := make(map[int]int)
	// 	for _,i := range input1 {
	// 		tempMap[i] += 1
	// 	}
	// 	for _,i := range input2 {
	// 		tempMap[i] += 1
	// 	}
	// 	for key, value := range tempMap {
	// 		if value > 1 {
	// 			tempInt = append(tempInt, key)
	// 		}
	// 	}
	// 	fmt.Println(tempInt)

}

func interceptionSet(inp1, inp2 []int)  {

	tempInt := make([]int, 0)
	for _, value1 := range inp1 {
		for _, value2 := range inp2 {
			if value1 == value2 {
				tempInt = append(tempInt, value1)
			}
		}
	}
	fmt.Printf("Результат пересечения двух множеств равен: %v \n", tempInt)
}
