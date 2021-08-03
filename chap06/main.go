package main

import (
	"fmt"
	"time"
)

// define a struct
type Employee struct {
	id int
	name string
	country string
	created time.Time
}

func main(){

	// declare variable
	var emp Employee
	newEmp := new(Employee)

	// set values
	emp.id = 2
	emp.name = "Jose Tandavala"
	emp.country = "UK"
	emp.created = time.Now()

	newEmp.id = 5
	newEmp.name = "Angelino Valeta"
	newEmp.country = "UK"
	newEmp.created = time.Now()

	// display
	fmt.Println("===================================")
	fmt.Println(emp.id)
	fmt.Println(emp.name)
	fmt.Println(emp.country)
	fmt.Println(emp.created)

	fmt.Println("=====================================")
	fmt.Println(newEmp.id)
	fmt.Println(newEmp.name)
	fmt.Println(newEmp.country)
	fmt.Println(newEmp.created)
}
