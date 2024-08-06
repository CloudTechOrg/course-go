package main

import (
	"fmt"
)

func main() {
	var scores []int

	scores = append(scores, 10)
	scores = append(scores, 20)
	scores = append(scores, 30)

	for _, score := range scores {
		fmt.Println(score)
	}

}
