package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (person Person) setValue(name string, age int) {
	person.Name = name
	person.Age = age
}

func main() {
	person := Person{
		Name: "テスト太郎",
		Age:  30,
	}

	fmt.Println("---初期の値を表示---")
	fmt.Println(person)

	person.setValue("テスト二郎", 25)

	fmt.Println("---メソッド実行後の値を表示---")
	fmt.Println(person)
}
