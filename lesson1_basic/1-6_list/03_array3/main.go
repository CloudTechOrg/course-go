package main

import (
	"fmt"
)

func main() {
	scores := [3]int{10, 20, 30}

	for i := 0; i < len(scores); i++ {
		fmt.Println(scores[i])
	}
}
