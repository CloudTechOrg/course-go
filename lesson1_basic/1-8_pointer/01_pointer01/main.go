package main

import "fmt"

func main() {
	var value1 int = 10
	var value2 int = value1

	// アドレスを出力してみる
	fmt.Println("---アドレスの出力---")
	fmt.Println("value1:", &value1)
	fmt.Println("value2:", &value2)

	// まずは現時点の値を出力
	fmt.Println("---初期状態の値---")
	fmt.Println("value1:", value1)
	fmt.Println("value2:", value2)

	// valueの内容を変更
	value1 = 20
	fmt.Println("---valueの変更後---")
	fmt.Println("value1:", value1)
	fmt.Println("value2:", value2)

	// pointerの内容を変更
	value2 = 30
	fmt.Println("---pointerの変更後---")
	fmt.Println("value1:", value1)
	fmt.Println("value2:", value2)
}
