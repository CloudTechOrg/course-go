package main

import (
	"fmt"
)

func main() {
	// Goで扱える変数の型

	//  int・・・整数値
	var int_value int = 10
	fmt.Println(int_value)

	//  String・・・文字列
	var string_value string = "Hello World"
	fmt.Println(string_value)

	//  float・・・小数
	//    float32・・・32ビット、約7桁の精度
	//    float64・・・64ビット、約15桁の精度
	var float_value float32 = 3.14
	fmt.Println(float_value)

	// bool・・・真偽（true or false）
	var bool_value bool = true
	fmt.Println(bool_value)
}
