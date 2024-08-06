package main

import (
	"fmt"
)

func main() {
	result := add(10, 20)
	fmt.Println(result)
}

func add(value1 int, value2 int) int {
	return value1 + value2
}
