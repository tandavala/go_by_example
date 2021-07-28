package main

import (
	"fmt"
)

func add(numbers ...int) int {
	result := 0
	for _, number := range numbers {
	result += number
	}
	return result
}

func main(){
	fmt.Println(add(12, 13, 14, 15))
}
