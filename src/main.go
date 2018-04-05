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

	TestCase4("2017-01-01", 1, 0, 1)
	TestCase4("2017-02-28", 1, 0, 1)
	TestCase4("2017-01-01", 2, 0, 1)
	TestCase4("2017-02-28", 2, 0, 1)
	fmt.Println("-----------")
	TestCase4("2017-01-01", 1, 0, 2)
	TestCase4("2017-02-28", 1, 0, 2)
	TestCase4("2017-01-01", 2, 0, 2)
	TestCase4("2017-02-28", 2, 0, 2)
	fmt.Println("-----------")
	TestCase4("2017-01-01", 1, 1, 1)
	TestCase4("2017-02-08", 1, 1, 1)
	TestCase4("2017-01-01", 1, 18, 1)
	TestCase4("2017-02-08", 1, 18, 1)
	fmt.Println("-----------")
	TestCase4("2017-01-01", 2, 1, 1)
	TestCase4("2017-02-08", 2, 1, 1)
	TestCase4("2017-01-01", 2, 18, 1)
	TestCase4("2017-02-08", 2, 18, 1)
	fmt.Println("-----------")
	TestCase4("2017-01-06", 2, 1, 1)
	TestCase4("2017-02-09", 2, 1, 1)
	TestCase4("2017-01-19", 2, 18, 1)
	TestCase4("2017-02-18", 2, 18, 1)
	fmt.Println("-----------")
	TestCase5("2017-01-06", 2, 1, 1, 2)
	TestCase5("2017-02-09", 2, 1, 1, 2)
	TestCase5("2017-01-19", 2, 18, 1, 2)
	TestCase5("2017-02-18", 2, 18, 1, 2)
	fmt.Println("-----------")
	TestCase5("2017-01-06", 2, 1, 1, 3)
	TestCase5("2017-02-09", 2, 1, 1, 3)
	TestCase5("2017-01-19", 2, 18, 1, 3)
	TestCase5("2017-02-18", 2, 18, 1, 3)
	fmt.Println("-----------")
	TestCase6("2017-01-06", 2, 1, 1, 12)
	TestCase6("2017-02-09", 2, 1, 1, 12)
	TestCase6("2017-01-19", 2, 18, 1, 12)
	TestCase6("2017-02-18", 2, 18, 1, 12)
	fmt.Println("-----------")
	TestCase5("2017-01-06", 0, 0, 1, 3)
	TestCase5("2017-02-09", 0, 0, 1, 3)
	TestCase5("2017-01-19", 0, 0, 1, 3)
	TestCase5("2017-02-18", 0, 0, 1, 3)

	fmt.Println("-------------------------------")

	TestRegExp()
	fmt.Println("-------------------------------")

	Interface2Struct()
	fmt.Println("-------------------------------")

	TestReflect()
	fmt.Println("-------------------------------")

	TestReflect2()
	fmt.Println("-------------------------------")

	TestPayLimit(7500.00, 2500.00)
	TestPayLimit(5000.00, 2000.00)
	TestPayLimit(3000.00, 2000.00)
	TestPayLimit(1000.00, 1000.00)
	TestPayLimit(500.00, 500.00)

	TestPayLimit(500.00, 200.00)
	TestPayLimit(500.00, 2500.00)
	TestPayLimit(2000.00, 2000.00)
	fmt.Println("-------------------------------")

	TestPayLimit2(7500.88, 2500.88)
	TestPayLimit2(5000.99, 2000.99)
	TestPayLimit2(3000.01, 2000.01)
	TestPayLimit2(1000.54, 1000.54)
	TestPayLimit2(500.45, 500.45)
	TestPayLimit2(2000.00, 2000.00)
	TestPayLimit2(2500.00, 2000.00)
	TestPayLimit2(2100.00, 2000.00)

	TestPayLimit2(2100.00, 2100.00)
	TestPayLimit2(500.00, 500.01)
	TestPayLimit2(500.00, 499.99)
	TestPayLimit2(500.99, 500.00)
	TestPayLimit2(3500.99, 2500.00)
	TestPayLimit2(2000.00, 1999.91)
	fmt.Println("xxxxxxxxxxxxxxxx1")
	TestPayLimit2(4000.00, 2000.00)
	TestPayLimit2(4000.00, 3000.00)
	TestPayLimit2(4000.00, 4000.00)
	fmt.Println("xxxxxxxxxxxxxxxx2")
	TestPayLimit2(4000.00, 1000.00)
	TestPayLimit2(4000.00, 2500.00)
	TestPayLimit2(4000.00, 3300.00)
	TestPayLimit2(4000.00, 5000.00)
	fmt.Println("xxxxxxxxxxxxxxxx3")
	TestPayLimit2(4000.54, 2000.54)
	TestPayLimit2(4000.45, 3000.45)
	TestPayLimit2(4000.99, 4000.99)
	TestPayLimit2(2890.99, 2890.99)
	TestPayLimit2(2890.99, 1890.99)
	TestPayLimit2(4000.54, 3000.454)
	TestPayLimit2(4000.54, 3000.544)
	fmt.Println("-------------------------------")

	TestPayLimit3(7500.88, 2500.88)
	TestPayLimit3(5000.99, 2000.99)
	TestPayLimit3(3000.01, 2000.01)
	TestPayLimit3(1000.54, 1000.54)
	TestPayLimit3(500.45, 500.45)
	TestPayLimit3(2000.00, 2000.00)
	TestPayLimit3(2500.00, 2000.00)
	TestPayLimit3(2100.00, 2000.00)
	TestPayLimit3(2100.00, 2100.00)
	TestPayLimit3(500.00, 500.01)
	TestPayLimit3(500.00, 499.99)
	TestPayLimit3(2000.00, 1999.91)
	TestPayLimit3(4000.00, 2000.00)
	TestPayLimit3(4000.00, 3000.00)
	TestPayLimit3(4000.00, 4000.00)
	TestPayLimit3(4000.00, 1000.00)
	TestPayLimit3(4000.00, 2500.00)
	TestPayLimit3(4000.00, 3300.00)
	TestPayLimit3(4000.00, 5000.00)
	TestPayLimit3(4000.54, 2000.54)
	TestPayLimit3(4000.45, 3000.45)
	TestPayLimit3(4000.99, 4000.99)
	TestPayLimit3(2890.99, 2890.99)
	TestPayLimit3(2890.99, 1890.99)
	TestPayLimit3(4000.54, 3000.454)
	TestPayLimit3(4000.54, 3000.544)
	fmt.Println("-------------------------------")

}
