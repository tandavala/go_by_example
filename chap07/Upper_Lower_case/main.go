package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println("----- demo upper/lower characters -------")
	message := "Hello World, GO!"
	fmt.Println(message)
	fmt.Println(strings.ToUpper(message))
	fmt.Println(strings.ToLower(message))
}
