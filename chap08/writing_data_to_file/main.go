package main

import (
	"fmt"
	"io/ioutil"
)

func writeFile(message string){
	bytes := []byte(message)
	ioutil.WriteFile("./testgo.txt", bytes, 0644)
	fmt.Println("created a file")
}


func main(){
	writeFile("Golang Developer, tandavala")
}

