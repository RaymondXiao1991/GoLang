package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TestPayLimit 支付限制
func TestPayLimit(shouldPay, payAmount float64) {
	if shouldPay > 2000.00 {
		if payAmount < 2000.00 {
			fmt.Println("不允许支付!")
		} else {
			fmt.Println("允许支付!")
		}
	} else {
		if shouldPay != payAmount {
			fmt.Println("不允许支付!")
		} else {
			fmt.Println("允许支付!")
		}
	}
}

// TestPayLimit2 支付限制
func TestPayLimit2(shouldPay, payAmount float64) {
	shouldPay = Precision(shouldPay, 2, true)
	//payAmount = Precision(payAmount, 2, true)

	if payAmount > shouldPay {
		fmt.Println("TEST CASE: shouldPay:", shouldPay, "payAmount:", payAmount, "支付金额错误,本次支付金额:", payAmount, "大于剩余可付金额:", shouldPay)
		return
	}

	minCanPay := shouldPay
	for minCanPay >= 3000 {
		minCanPay -= 1000
	}
	minCanPay = Precision(minCanPay, 2, true)
	if payAmount < minCanPay {
		fmt.Println("TEST CASE: shouldPay:", shouldPay, "payAmount:", payAmount, "支付金额错误,剩余可付金额:", shouldPay, ",请至少支付:", minCanPay)
	} else {
		if int64((payAmount*1000-minCanPay*1000)/1000)%1000 != 0 {
			fmt.Println("判断:", payAmount*1000-minCanPay*1000)
			fmt.Println("判断:", (payAmount*1000-minCanPay*1000)/1000)
			fmt.Println("判断:", int64((payAmount*1000-minCanPay*1000)/1000)%1000)
			fmt.Println("TEST CASE: shouldPay:", shouldPay, "payAmount:", payAmount, "支付金额错误,剩余可付金额:", shouldPay, ",请至少支付:", minCanPay, ",并每次加1000")
		} else {
			fmt.Println("判断:", payAmount*1000-minCanPay*1000)
			fmt.Println("判断:", (payAmount*1000-minCanPay*1000)/1000)
			fmt.Println("判断:", int64((payAmount*1000-minCanPay*1000)/1000)%1000)
			fmt.Println("TEST CASE: shouldPay:", shouldPay, "payAmount:", payAmount, "允许支付!")
		}
	}
}

// PayLimit 支付限制
func PayLimit(shouldPay, payAmount float64) error {
	canPay := shouldPay
	for canPay > 2000 {
		if canPay-1000 < 2000 {
			break
		}
		canPay -= 1000
	}
	canPay = Precision(canPay, 2, true)
	if payAmount < canPay {
		return fmt.Errorf("支付金额错误,剩余可付金额:%f,请至少支付:%f", shouldPay, canPay)
	} else if int64(payAmount-canPay)%1000 != 0 {
		return fmt.Errorf("支付金额错误,剩余可付金额:%f,请至少支付:%f,并每次加1000", shouldPay, canPay)
	}

	return nil
}

// TestPayLimit3 支付限制
func TestPayLimit3(shouldPay, payAmount float64) {
	shouldPay = Precision(shouldPay, 2, true)
	//payAmount = Precision(payAmount, 2, true)
	/*
		if payAmount > shouldPay {
			fmt.Println("TEST CASE: shouldPay:", shouldPay, "payAmount:", payAmount, "支付金额错误,本次支付金额:", payAmount, "大于剩余可付金额:", shouldPay)
			return
		}
		else if payAmount == shouldPay {
			fmt.Println("TEST CASE: shouldPay:", shouldPay, "payAmount:", payAmount, "允许支付!")
			return
		}
	*/

	minCanPay := shouldPay
	canPaySlice := []string{}
	canPaySlice = append(canPaySlice, strconv.FormatFloat(shouldPay, 'f', 2, 64))
	for minCanPay >= 3000 {
		minCanPay = Precision(minCanPay-1000, 2, true)
		canPaySlice = append(canPaySlice, strconv.FormatFloat(minCanPay, 'f', 2, 64))
	}
	fmt.Println("canPaySlice:", canPaySlice)

	if !SliceContains(canPaySlice, Format2String2(payAmount)) { // 小数位为零时丢失有效位数
		//if !SliceContains(canPaySlice, strconv.FormatFloat(payAmount, 'f', 2, 64)) { // 指定两位有效小数
		fmt.Println("TEST CASE: shouldPay:", shouldPay, "payAmount:", payAmount, "支付金额错误,剩余可付金额:", shouldPay, ",请至少支付:", minCanPay, ",并每次加1000")
	} else {
		fmt.Println("TEST CASE: shouldPay:", shouldPay, "payAmount:", payAmount, "允许支付!")
	}
}

// SliceContains 判断是否包含
func SliceContains(src []string, value string) bool {
	isContain := false
	for _, v := range src {
		if v == value {
			isContain = true
			break
		}
	}
	return isContain
}

// Format2String 格式化
func Format2String(f float64) string {
	// 1.强转至int64,再转回float64,判断是否等于本身
	/*
		if float64(int64(f)) == f {
			return strconv.FormatFloat(f, 'f', 2, 64)
		}
	*/

	// 2.转为string,再判断是否有"."
	str := strconv.FormatFloat(f, 'f', -1, 64)

	if len(strings.Split(str, ".")[1:]) < 2 {
		return strconv.FormatFloat(f, 'f', 2, 64)
	}

	return str

	// 3.乘以一个大数,再求余数是否等于零
}

// Format2String2 格式化
func Format2String2(f float64) string {
	return fmt.Sprintf("%.2f", f)
}

// TestFor
func TestFor() {
	checkTerm := 0
	hasTerms := 1
	fmt.Println("[应出账期数:", checkTerm, "],", "[已出账期数:", hasTerms, "]")
	for currentTerm := hasTerms + 1; currentTerm <= checkTerm; currentTerm++ {
		fmt.Println("[开始生成第", currentTerm, "期账单]")
	}
}
