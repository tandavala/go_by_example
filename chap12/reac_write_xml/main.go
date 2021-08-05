package main

import (
	"fmt"
	"encoding/xml"
)

func main(){
	demoXml()
}

func demoXml(){
	fmt.Println("-------- Demo encoding xml --------")
	
	type EmployeeCountry struct {
		CountryCode string `xml:"code;attr"` // XML attribute: code
		CountryName string `xml:",chardata" // XML value`
	}
	
	type Employee struct {
		XMLName xml.Name `xml:"employee"`
		Id string `xml:"id"`
		Name string `xml:"name"`
		Email string `xml:"email"`
		Country EmployeeCountry `xml:"country"`
	}

	// struct to xml
	fmt.Println(">>>>>struct to xml.....")
	emp := &Employee{Id:"12345",Name:"Tandavala",Email:"jose.tandavala@gmail.com",
					 Country: EmployeeCountry{CountryCode:"AO",CountryName:"Angola"}}
	b, err := xml.Marshal(emp)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	// xml string to struct
	fmt.Println(">>>>>>xml to struct.....")
	var newEmp Employee
	str := `<employee><id>555</id><name>Jose</name><email>jose@gmail.com</email>`
	xml.Unmarshal([]byte(str),&newEmp)
	fmt.Printf("Id: %s\n", newEmp.Id)
	fmt.Printf("Name: %s\n", newEmp.Name)
	fmt.Printf("Email: %s\n", newEmp.Email)

}
