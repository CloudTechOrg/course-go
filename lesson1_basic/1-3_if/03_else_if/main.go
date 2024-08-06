package main

import "fmt"

func main() {
	score := 59

	// --------------------
	// 条件分岐（else if）
	// --------------------
	// 条件をより細かく定義したいときに利用

	// 複数の条件を定義し、該当すればそこで条件分岐の処理を終了
	// どれにも該当しない場合、elseの処理を実施
	if score >= 80 {
		fmt.Println("大変良くできました")
	} else if score >= 60 {
		fmt.Println("あと一息です")
	} else {
		fmt.Println("もっと頑張りましょう")
	}
}
