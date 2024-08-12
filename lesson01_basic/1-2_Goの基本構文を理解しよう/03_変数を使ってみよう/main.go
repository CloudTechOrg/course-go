package main

import (
	"fmt"
)

func main() {
	// 変数は、文字列や数値などの値を格納する"箱"のようなもの
	// 変数に値を入れることで、それを使い回すことが出来る

	// 先程の四則演算の処理に当てはめてみる
	// 変数の宣言と、値の格納
	// var 変数名 型 = 値
	var value1 int = 20
	var value2 int = 10

	// 計算結果の表示
	fmt.Println(value1 + value2)

	// 引き算
	fmt.Println(value1 - value2)

	// 掛け算
	fmt.Println(value1 * value2)

	// 割り算
	fmt.Println(value1 / value2)
}
