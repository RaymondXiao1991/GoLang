package main

import (
	"fmt"
	"time"
)

func main() {

	// 测试时间差计算
	TestCaseOfCalcTime()
	fmt.Println("-------------------------------")

	// 测试今日的七天后
	sevenDaysLater := time.Now().AddDate(0, 0, 7)
	fmt.Println("sevenDaysLater:", sevenDaysLater)
	fmt.Println("-------------------------------")

	totalTerm, _ := CalcTotalBillTerm(2, 6)
	fmt.Println("totalTerm:", totalTerm)
	fmt.Println("-------------------------------")

	payableAmount := Precision(100.1234, 2, true)
	fmt.Println("payableAmount:%f", -payableAmount)
	shouldPay := fmt.Sprintf("￥%.2f", -payableAmount) // 应付总额
	fmt.Println("shouldPay:", shouldPay)
	fmt.Println("-------------------------------")

	fmt.Println(GetLaterDate(time.Now(), 0))
	fmt.Println("-------------------------------")

	var payable, paid float64
	payable = 292.600000
	paid = 0.00000
	shouldPayAmount := (payable*1000 - paid*1000) / 1000
	fmt.Println("shouldPayAmount:", shouldPayAmount)
	fmt.Println("-------------------------------")

	mapPriorities()
}
