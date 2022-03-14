package main

import "fmt"

func main() {
	// Создаем временную мапу, в которой строки статут ключами, затем ключи добавим в список, тем самым получим собственое множество
	tempMap := make(map[string]struct{})
	list := []string{"cat", "cat", "dog", "cat", "tree"} 
	set := make([]string, 0)

	for _, i := range list {
		tempMap[i] = struct{}{}
	}

	for key := range tempMap {
		set = append(set, key)
	}

	fmt.Println(set)
}