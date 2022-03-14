package main

import "fmt"

func main() {
	vStr := "hello world"
	vInt := 123
	vBool := true
	vChan := make(chan int)
	vList := make([]interface{}, 0)
	vList = append(vList, vStr, vInt, vBool, vChan)
	for _,i := range vList {
		fmt.Println(typeof(i))
	}
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}