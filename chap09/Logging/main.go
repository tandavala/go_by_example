package main

import (
	"fmt"
	"log"
	//"os"
)

func main(){
	simpleLogging()
}

func simpleLogging(){
	fmt.Println("------ simple logging ------")
	log.Println("Hello World")
	log.Println("this is a simple error")
}
