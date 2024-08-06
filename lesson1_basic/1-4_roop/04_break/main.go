package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)

		if i == 5 {
			// ここでループを抜ける
			break
		}
	}

	fmt.Println("ループを終了しました")

}
