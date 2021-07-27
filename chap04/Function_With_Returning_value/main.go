package main

import (
	"fmt"
)

func advanced_calculate(a, b, c float64) float64 {
	result := a*b*c
	return result
}

func main(){
	fmt.Println(advanced_calculate(1,2,3))
}
