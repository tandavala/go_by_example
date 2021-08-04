package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	//simpleLogging()
	formattingLogging()
}

func simpleLogging(){
	fmt.Println("------ simple logging ------")
	log.Println("Hello World")
	log.Println("this is a simple error")
}

func formattingLogging(){
	fmt.Println("--------- formattingLogging---------")
	var warning *log.Logger

	warning = log.New(
		os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	warning.Println("This is warning message 1")
	warning.Println("This warning message 2")
}
