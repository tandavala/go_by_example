package main

import (
	"fmt"
	//"strconv"
	//"strings"
)

func main(){
	str1 := "hello"
	str2 := "world"
	str3 := str1 + str2
	fmt.Println(str3)

	str4 := fmt.Sprintf("%s %s", str1, str2)
	fmt.Println(str4)
}
