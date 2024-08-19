package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "テスト太郎", Age: 30}

	fmt.Println(p.Name)
	fmt.Println(p.Age)
}
