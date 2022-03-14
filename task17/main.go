package main

import (
	"fmt"
	"sort"
)

// Бинарный поиск числа
func SearchInIntSlice(haystack []int, needle int) (result bool, iterationsCount int) {
	// Данный алгоритм работает только с отсартированными массивами.
    sort.Ints(haystack) 
    lowKey := 0 // первый индекс
    highKey := len(haystack) - 1 // последний индекс
    if (haystack[lowKey] > needle) || (haystack[highKey] < needle) {
        return // нужное значение не в диапазоне данных
    }
    for lowKey <= highKey { 
        // уменьшаем список рекурсивно
        iterationsCount++
        mid := (lowKey + highKey) / 2 // середина
        if haystack[mid] == needle {
            result = true // мы нашли значение
            return
        }
        if haystack[mid] < needle { 
            // если поиск больше середины - мы берем только блок с большими значениями увеличивая lowKey
            lowKey = mid + 1
            continue
        }
        if haystack[mid] > needle { 
            // если поиск меньше середины - мы берем блок с меньшими значениями уменьшая highKey
            highKey = mid - 1
        }
    }
    return
}
func main() {
	ar := []int{3, 4, 1, 2, 5, 7, -1, 0}
	num := 1
	ok,i := SearchInIntSlice(ar, num)

	if ok {
		fmt.Printf("Значение найдено за %v итерации \n", i)
	} else {
		fmt.Println("Число не найдено в массиве")
	}
	
}

