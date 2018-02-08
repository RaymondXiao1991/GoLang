package main

import (
	"fmt"
	"math"
)

// 控制小数位
func Precision(f float64, prec int, round bool) float64 {
	pow10_n := math.Pow10(prec)
	if round {
		if f >= 0 {
			return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
		}
		return math.Trunc((f-0.5/pow10_n)*pow10_n) / pow10_n
	}
	return math.Trunc((f)*pow10_n) / pow10_n
}

func pre() {
	fmt.Println(0.31 / 31 * (-26))
	recordServiceFee := Precision((0.31 / 31 * float64(-26)), 2, true)
	fmt.Println(recordServiceFee)
	fmt.Println("-------------------------------")

	fmt.Println(0.31 / 31 * (26))
	recordServiceFee = Precision((0.31 / 31 * float64(26)), 2, true)
	fmt.Println(recordServiceFee)
	fmt.Println("-------------------------------")

	var fa float64
	fa = 4964.11
	amount := int(fa*1000) / 10
	fmt.Println("amount:", amount)
	fmt.Println("-------------------------------")
}
