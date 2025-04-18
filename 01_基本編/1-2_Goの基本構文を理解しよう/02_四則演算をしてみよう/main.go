package main

import "fmt"

func main() {

	// -----------------------------
	// 基本的な四則演算
	// -----------------------------

	// 足し算は、+ の記号を使用する
	// %d は、整数値を文字列に埋め込むために使用される"
	fmt.Println(20 + 10)

	// 引き算は、- の記号を使用する
	fmt.Println(20 - 10)

	// 掛け算は、* の記号を使用する
	fmt.Println(20 * 10)

	// 割り算は、/ の記号を使用する
	fmt.Println(20 / 10)

	// -----------------------------
	// 割り算に関する補足
	// -----------------------------

	// 割り算の余りは、% の記号で求められる
	fmt.Println(20 % 7)

	// ゼロで割り算をしようとするとエラーになる
	// // はコメントと言い、実行されない処理を表す
	// // を外すと、「invalid operation: division by zero」のエラーになる
	// fmt.Println(10 / 0)
}
