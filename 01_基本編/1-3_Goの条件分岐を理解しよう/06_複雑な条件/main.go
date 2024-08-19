package main

import "fmt"

func main() {
	var value int

	// --------------------
	// 条件分岐（複雑な条件）
	// --------------------

	value = 100

	// かつ（And）
	// &&で条件を結合すると、「両方の条件に一致する場合」となる
	if value > 20 && value < 30 {
		fmt.Println("AND条件に一致します")
	}

	// または（Or）
	// ||で条件を結合すると、「いずれかの条件に一致する場合」となる
	if value == 10 || value == 20 {
		fmt.Println("OR条件に一致します")
	}

	// 否定
	// ! を記載すると、条件の否定になります。
	if value != 10 {
		fmt.Println("否定の条件に一致します")
	}

	// 条件の優先順位
	// if 条件1 || 条件2 && 条件3

}
