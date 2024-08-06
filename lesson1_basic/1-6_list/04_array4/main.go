package main

import (
	"fmt"
)

func main() {
	scores := [3]int{10, 20, 30}

	for i, score := range scores {
		fmt.Println(i)
		fmt.Println(score)
	}
}
