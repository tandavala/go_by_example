package main

import (
	"fmt"
	"encoding/json"
)

func main(){
	demoJson()
}

func demoJson(){
	fmt.Println("------- Demo encoding json --------")
	type Employee struct {
		Id string `json:"id"`
		Name string `json:"name"`
		Email string `json:"email"`
	}
	// struct to json
	fmt.Println(">>>> struct to json....")
	emp := &Employee{Id:"12345", Name:"Tandavala",Email:"jose.tandavala@ideiasdinamicas.com"}
	b, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	// json string to struct
	fmt.Println(">>>>>>>json to struct......")
	var newEmp Employee
	str := `{"Id":"4567","Name":"Valeta","Email":"angelino.valeta@ideiasdinamicas.com"}`
	json.Unmarshal([]byte(str), &newEmp)
	fmt.Printf("Id: %s\n", newEmp.Id)
	fmt.Printf("Name: %s\n", newEmp.Name)
	fmt.Printf("Email: %s\n", newEmp.Email)
}
