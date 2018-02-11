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
	fmt.Println("-------------------------------")

	fmt.Println(IntVector{3, 4, 5}.Sum())
	fmt.Println("-------------------------------")

	pointer()
	fmt.Println("-------------------------------")

	StartEndTime()
	EndDateOfThisMonth()
	fmt.Println("-------------------------------")

	fmt.Println("GetEndDateOfCurrentMonth:", GetEndDateOfCurrentMonth())
	fmt.Println("-------------------------------")

	fmt.Println("GetEndDateOfCurrentMonth:", GetEndDateOfCurrentMonth2(1506182400))
	fmt.Println("-------------------------------")

	startDateOfNextMonth := GetStartDateOfNextMonth(1517414400)
	fmt.Println("GetStartDateOfNextMonth:", startDateOfNextMonth)
	fmt.Println("-------------------------------")
	currentTermEndTime := AddMonthTime(time.Unix(startDateOfNextMonth.Unix(), 0), 1)
	fmt.Println("currentTermEndTime:", currentTermEndTime)
	fmt.Println("-------------------------------")

	today := time.Unix(GetLaterDateTime(time.Now(), 0), 0)
	fmt.Println("today:", today)
	fmt.Println("-------------------------------")

	testIf()
	fmt.Println("-------------------------------")

	TestAddMonth()
	fmt.Println("-------------------------------")

	TestDateUntil()
	fmt.Println("-------------------------------")

	TestAccuracy()
	fmt.Println("-------------------------------")

	TestCalTermEndTimeStr()
	fmt.Println("-------------------------------")

	output := new(Output)
	output.BillCode = "xxx"
	SortAttribute(output)
	fmt.Println("-------------------------------")

	outputdetail := new(OutputDetail)
	outputdetail.BillCode = "xxx"
	outputdetail.BillName = "yyy"
	SortAttribute(outputdetail)
	fmt.Println("-------------------------------")

	TestIsNormal()
	fmt.Println("-------------------------------")

	fmt.Println("H" + "i")
	fmt.Println('H' + 'i')
	fmt.Println("-------------------------------")

	TestOrderRecord()
	fmt.Println("-------------------------------")
}
