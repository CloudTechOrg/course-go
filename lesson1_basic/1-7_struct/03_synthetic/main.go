package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	MainJob Job
}

type Job struct {
	Company string
	Type    string
}

func main() {

	person := Person{
		Name: "テスト太郎",
		Age:  30,
		MainJob: Job{
			Company: "テスト株式会社",
			Type:    "ITエンジニア",
		},
	}

	fmt.Println(person.Name)
	fmt.Println(person.Age)
	fmt.Println(person.MainJob.Company)
	fmt.Println(person.MainJob.Type)

}
