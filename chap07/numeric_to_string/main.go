package main

import (
	"fmt"
//	"strings"
)

func main(){
	fmt.Println("------- demoe numeric to string -------")
	num1 := 8
	num2 := 5.872

	var str_num1 string
	str_num1  = fmt.Sprintf("%d", num1)
	fmt.Println(str_num1)

	var str_num2 string
	str_num2 = fmt.Sprintf("%f", num2)
	fmt.Println(str_num2)
}
