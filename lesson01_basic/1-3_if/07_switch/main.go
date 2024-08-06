package main

import "fmt"

func main() {
	value := 1

	// --------------------
	// 条件分岐（switch）
	// --------------------

	// ある変数の値の内容に応じて処理を振り分けたいときの書き方
	// ifとelse ifで同様の事ができるが、より簡潔に記載が出来る
	switch value {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Other")
	}

}
