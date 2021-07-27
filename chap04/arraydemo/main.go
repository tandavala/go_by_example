package main

import (
	"fmt"
	"math"
)

func main(){
	foo()
	circle_area(2.8)
}

// a simple function
func foo(){
	fmt.Println("foo() was called")
}

func circle_area(r float64){
	area := math.Pi * math.Pow(r, 2)
	fmt.Printf("Circle area (r=%.2f) = %.2f \n", r, area)
}
