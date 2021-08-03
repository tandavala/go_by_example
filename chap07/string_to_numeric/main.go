package main

import (
	"fmt"
	"strconv"
	//"strings"
)

func main(){
	fmt.Println("------------ demoe String To Numeric-----------")
	str_val1 := "5"
	str_val2 := "2.8769"

	var err error
	var int_val1 int64
	int_val1, err = strconv.ParseInt(str_val1, 10, 32) // base10, 32 size
	if err== nil {
		fmt.Println(int_val1)
	} else {
		fmt.Println(err)
	}
	var float_val2 float64
	float_val2, err = strconv.ParseFloat(str_val2, 64) // 64 size
	if err==nil {
		fmt.Println(float_val2)
	} else {
		fmt.Println(err)
	}
}
