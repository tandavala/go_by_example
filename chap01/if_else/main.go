package main

import "fmt"

func main(){
	var (
		a = 5
		b = 8
 )
	if a > b || a - b < a {
	fmt.Println("Conditional --> a>b || a - b < a")
	} else {
		fmt.Println("...another")
	}
}
