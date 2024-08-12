package main

import "fmt"

func test_func(value int) {
	value = 100
}

func main() {
	// 整数変数の定義
	var value int = 10

	fmt.Println("---初期の値を表示---")
	fmt.Println("value:", value)

	test_func(value)

	fmt.Println("---関数実行後の値を出力---")
	fmt.Println("value:", value)
}
