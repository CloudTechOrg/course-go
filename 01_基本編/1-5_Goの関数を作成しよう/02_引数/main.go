package main

import (
	"fmt"
)

func main() {
	add(10, 20)
}

func add(value1 int, value2 int) {
	result := value1 + value2
	fmt.Println(result)
}
