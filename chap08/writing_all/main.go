package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	// Write data into a file
	fmt.Println("Writing data into a file")
	writeFile("Welcome to go")

	//( read data from a file
	fmt.Println("reading data from a file")
	readFile()
}

func writeFile(message string){
	bytes := []byte(message)
	ioutil.WriteFile("./testgo.txt", bytes, 0644)
	fmt.Println("created file")
}

func readFile(){
	data, _ := ioutil.ReadFile("./testgo.txt")
	fmt.Println("file content")
	fmt.Println(string(data))
}
