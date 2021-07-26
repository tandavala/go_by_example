package main

import "fmt"

func main(){
// declare variables
var a, b int

// assing values
a = 5
b = 10

// arithmetic operation
// additon
c := a + b
fmt.Println("%d + %d = %d \n", a, b, c)

// Subtraction
d := a - b
fmt.Println("%d - %d = %d \n", a, b, d)

// division
e :=  float32(a) / float32(b)
fmt.Println("%d / %d = %d \n", a, b, e)

// multiplication
f := a * b
fmt.Println("%d * %d = %d \n", a, b, f)

}
