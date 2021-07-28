package main

import (
	"fmt"
)

func compute(a, b, c float64, name string) (float64, float64, string) {
	result1 := a*b*c
	result2 := a+b+c
	result3 := result1 + result2
	newName := fmt.Sprintf("%s value = %.2f", name, result3)

	return result1, result2, newName
}

func main(){
	fmt.Println(compute(16, 17, 20, "tandavala"))
}
