package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println("------ demoe String Parser -----")
	data := "Berlin;Amsterdam;London;Tokyo"
	fmt.Println(data)
	cities := strings.Split(data, ";")
	fmt.Println(cities)
	for _, city := range cities {
		fmt.Println(city)
	}
}
