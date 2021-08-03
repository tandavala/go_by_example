package main

import (
	"fmt"
	"io/ioutil"
)

func readFile(){
	data, _ := ioutil.ReadFile("./testgo.txt")
	fmt.Println("file content: ")
	fmt.Println(string(data))
}

func main(){
	readFile()
}
