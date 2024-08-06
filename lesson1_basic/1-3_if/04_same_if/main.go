package main

import "fmt"

func main() {

	// 「条件に一致する場合」は、= ではなく == とする必要がある
	value := 10
	if value == 10 {
		fmt.Println("1つ目の条件に一致します")
	}

	// =を条件として記載すると構文エラーになる
	// if value = 10 {
	// 	fmt.Println("条件に一致します")
	// }

}
