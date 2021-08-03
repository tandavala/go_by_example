package main

import "fmt"

func main(){
	var x int

	x = 10
	fmt.Println(x)
	fmt.Println(&x)

	// declare pointer
	var num *int
	val := new(int)

	num = new(int)
	*num = x // set value

	val = &x // set address

	fmt.Println("========= pointer num ============")
	fmt.Println(*num) // print a value of x
	fmt.Println(num) // print address
	fmt.Println("================  pointer val =======")
	fmt.Println(*val)
	fmt.Println(val) // print address of x

}
