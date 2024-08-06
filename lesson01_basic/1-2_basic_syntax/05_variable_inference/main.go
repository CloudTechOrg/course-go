package main

import (
	"fmt"
)

func main() {
	// 型推論
	// 型を定義しなくても、値の内容を推論して型を定義してくれる

	// 整数値を推論
	result1 := 10
	fmt.Println(result1)

	// 小数と整数を計算すると、小数になる
	result2 := 3 + 0.14
	fmt.Println(result2)

	// 整数値同士の割り算は、整数になる
	result3 := 10 / 3
	fmt.Println(result3)

	// 片方が小数なら、結果も少数になる
	result4 := 10 / 3.0
	fmt.Println(result4)

	// 他にも様々なパターンがあるので、色々試してみましょう。

}
