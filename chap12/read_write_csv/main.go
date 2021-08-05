package main

import (
	"fmt"
	"encoding/csv"
	"os"
)

func main(){
	demoCsv()
}

func demoCsv(){
	fmt.Println("-------- Demo encoding csv ----------")
	type Employee struct {
		Name string
		Email string
		Country string
	}

	// read csv to a array of struct
	fmt.Println(">>>>>>read a csv file and load to array of struct....")
	file, err := os.Open("./demo.csv")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3
	reader.Comma = ';'

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var emp Employee
	var employees []Employee

	for _, item := range csvData {
		emp.Name = item[0]
		emp.Email = item[1]
		emp.Country = item[2]
		employees = append(employees, emp)
		fmt.Printf("name %s email: %s country: %s\n", item[0], item[1], item[2])
	}
	fmt.Println(">>>>>show all employee structs.....")
	fmt.Println(employees)

	// write data to csv file
	fmt.Println(">>>>>>write data to a csv file.....")
	csvFile, err := os.Create("./demo2.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	writer.Comma = ';'
	for _, itemEmp := range employees {
		records := []string{itemEmp.Name,itemEmp.Email,itemEmp.Country}
		err := writer.Write(records)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	writer.Flush()
	fmt.Println("Done")
}
