package main

import (
	"fmt"
	"math/rand"
)

func main(){

	// define array
	fmt.Println("define map")
	products := make(map[string]int)
	customers := make(map[string]int)

	// insert data
	fmt.Println(">>>>>>> insert map data")
	products["Product1"] = rand.Intn(100)
	products["Product2"] = rand.Intn(100)
	products["Product3"] = rand.Intn(100)

	customers["cust1"] = rand.Intn(100)
	customers["cust2"] = rand.Intn(100)

	// display data
	fmt.Println(">>>>>>>> display map data")
	fmt.Println(products["Product1"])
	fmt.Println(products["Product2"])
	fmt.Println(products["Product3"])
	fmt.Println(customers["cust1"])
	fmt.Println(customers["cust2"])
}
