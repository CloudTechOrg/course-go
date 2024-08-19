package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	CompanyName string
	Department  string
}

func main() {

	employee := Employee{
		Person: Person{
			Name: "テスト太郎",
			Age:  30,
		},
		CompanyName: "テスト株式会社",
		Department:  "開発部",
	}

	fmt.Println(employee.Name)
	fmt.Println(employee.Age)
	fmt.Println(employee.CompanyName)
	fmt.Println(employee.Department)

}
