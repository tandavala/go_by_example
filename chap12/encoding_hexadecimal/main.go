package main

import (
	"fmt"
	"encoding/hex"
)

func main(){
	message := "hello, go (*w3hu%#"
	demoHex(message)
}

func demoHex(message string){
	fmt.Println("---------- Demo encoding Hex ---------")

	fmt.Println("plaintext: ")
	fmt.Println(message)

	encoding := hex.EncodeToString([]byte(message))
	fmt.Println("Hex msg:")
	fmt.Println(encoding)

	decoding, _ := hex.DecodeString(encoding)
	fmt.Println("decoding Hex msg:")
	fmt.Println(string(decoding))
}
