package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	fileLogging()
}

func fileLogging(){
	fmt.Println("--------- file logging ----------")
	file, err := os.OpenFile("./app.log",
			os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file")
		return
	}
	var logFile *log.Logger
	logFile = log.New(
		file,
		"APP: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	logFile.Println("This is error message 1")
	logFile.Println("This is error message 2")
	logFile.Println("This is error message 3")
}
