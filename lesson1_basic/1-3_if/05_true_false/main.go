package main

import "fmt"

func main() {

	// trueであれば、必ず条件に一致する
	if true {
		fmt.Println("1つ目の条件に一致します")
	}

	// falseであれば、必ず条件に一致しない
	if false {
		fmt.Println("2つ目の条件に一致します")
	}

	// bool型の変数でも表せる
	var value1 bool = true
	if value1 {
		fmt.Println("3つ目の条件に一致します")
	}

	// 条件自体を変数に格納できる
	var value2 bool = 10 == 10
	if value2 {
		fmt.Println("3つ目の条件に一致します")
	}

}
