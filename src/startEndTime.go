package main

import (
	"fmt"
	"time"
)

const DATE_FORMAT = "2006-01-02"

func StartEndTime() {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	start := thisMonth.AddDate(0, -1, 0).Format(DATE_FORMAT)
	end := thisMonth.AddDate(0, 0, -1).Format(DATE_FORMAT)
	timeRange := fmt.Sprintf("%s~%s", start, end)
	fmt.Println(timeRange)

	fmt.Println(thisMonth)
}

func EndDateOfThisMonth() {
	now := time.Now()
	year, month, _ := now.Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	start := thisMonth.AddDate(0, 0, 0).Format(DATE_FORMAT)
	end := thisMonth.AddDate(0, 1, -1).Format(DATE_FORMAT)
	timeRange := fmt.Sprintf("%s~%s", start, end)
	fmt.Println(timeRange)
}

// 获取本月末日期
func GetEndDateOfCurrentMonth() time.Time {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	return thisMonth.AddDate(0, 1, -1)
}

// 获取本月末日期
func GetEndDateOfCurrentMonth2(now int64) time.Time {
	year, month, _ := time.Unix(now, 0).Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	return thisMonth.AddDate(0, 1, -1)
}

// 获取下月初日期
func GetStartDateOfNextMonth(now int64) time.Time {
	year, month, _ := time.Unix(now, 0).Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	return thisMonth.AddDate(0, 1, 0)
}

//平方
type myint int

func (p myint) mysquare() int {
	p = p * p
	fmt.Println("mysquare p = ", p)
	return 0
}

// 获取本月末日期
type T struct {
	tt time.Time
}

func (now T) GetEndDateOfCurrentMonth2() time.Time {
	year, month, _ := now.tt.Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	return thisMonth.AddDate(0, 1, -1)
}

// 获取本月末日期
type INT int

func (now INT) GetEndDateOfCurrentMonth2() time.Time {
	year, month, _ := time.Unix(int64(now), 0).Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	return thisMonth.AddDate(0, 1, -1)
}
