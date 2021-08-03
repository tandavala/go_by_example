package main

import (
	"fmt"
)

func main(){
	fmt.Println("------- demoe coy data --------")
	sample := "hello world, go"
	fmt.Println(sample)
	fmt.Println(sample[0:len(sample)]) // copy all
	fmt.Println(sample[:5]) // copy 5 characters
	fmt.Println(sample[3:8]) // copy characters from index 3 until 8
	fmt.Println(sample[len(sample)-5:len(sample)]) // 5 copy characters
}
