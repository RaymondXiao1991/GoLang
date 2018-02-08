package main

import (
	"fmt"
	"math"
	"time"
)

// 获取几天后(精度只到天的时间)
func GetLaterDateTime(t time.Time, d int) int64 {
	return time.Date(t.Year(), t.Month(), t.Day()+d, 0, 0, 0, 0, time.Local).Unix()
}

func cal() {
	var a uint = 60
	var b uint = 13
	var c uint = 0

	c = a & b
	fmt.Printf("a & b = %d\n", c)

	c = a | b
	fmt.Printf("a | b = %d\n", c)

	c = a ^ b
	fmt.Printf("a ^ b = %d\n", c)

	c = a << 2
	fmt.Printf("a << 2 = %d\n", c)

	c = a >> 2
	fmt.Printf("a >> 2 = %d\n", c)

	var totalAmount float64 = 0.0100
	var alreadyPaid float64 = 0.0000
	var payAmount float64 = 0.01

	if totalAmount > (alreadyPaid*100000+payAmount*100000)/100000 {
		fmt.Println("xxxxxxxxxxxxxx")
	} else if totalAmount == (alreadyPaid*100000+payAmount*100000)/100000 {
		fmt.Println("yyyyyyyyyyyyyyyyyyy")
	} else {
		fmt.Printf("付款金额有误，超出账单总金额")
	}

	var shouldPay float64 = 0.01
	var TotalAmount float64 = 0.01
	shouldPay += +TotalAmount
	fmt.Println("shouldPay:", shouldPay)

	TotalAmount = 6000.00
	shouldPay = 5
	if payAmount > 0 && payAmount <= shouldPay {
		if (TotalAmount*100)/100 > 5000 && (payAmount*100)/100 >= 1000 && (payAmount*100)/100 <= 5000 { // 当应付金额>5000时，允许用户部分支付
			fmt.Println("true")
		} else if (TotalAmount*100)/100 < 5000 && payAmount == shouldPay {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	} else {
		fmt.Println("超过了需要支付的金额哦~")
	}
	if payAmount > 0 && payAmount <= shouldPay {
		if (TotalAmount*100)/100 > 5000 && (payAmount*100)/100 >= 1000 && (payAmount*100)/100 <= 5000 { // 当应付金额>5000时，允许用户部分支付
			fmt.Println("true")
		} else if (TotalAmount*100)/100 < 5000 && payAmount == shouldPay {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	} else {
		fmt.Println("超过了需要支付的金额哦~")
	}

	fmt.Println(Precision(5000-5.01, 2, true))
	fmt.Println(Precision(5000-5.01, 2, false))

	var date_time string
	date_time = "20175239"
	fmt.Println(date_time[0:4])
	fmt.Println(date_time[4:6])
	fmt.Println(date_time[6:8])

	str2 := "hello"
	data2 := []byte(str2)
	fmt.Println(data2)
	str2 = string(data2[:])
	fmt.Println(str2)
	str3 := string(data2)
	fmt.Println(str3)

	now := time.Now()
	endDate := time.Unix(1515254400, 0)
	nextDayTime := time.Unix(GetLaterDateTime(now, 1), 0)
	alreadyPayDiff := int(math.Ceil(nextDayTime.Sub(endDate.Add(time.Hour*24)).Hours() / 24.0))
	fmt.Println("alreadyPayDiff:", alreadyPayDiff)

	fmt.Println(Precision(5000-5.01, 2, true))
	fmt.Println(Precision(5000-5.01, 2, false))
}
