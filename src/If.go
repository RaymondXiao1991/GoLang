package main

import "fmt"

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func testIf() {
	a, b := 2, 3
	max := If(a > b, a, b).(int)
	fmt.Println("[", a, b, "] 大者为：", max)
}
