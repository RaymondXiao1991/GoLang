package main

import (
	"fmt"
	"time"
)

// AddMonthTime 获取几个月后的日期
func AddMonthTime(t time.Time, month int) time.Time {
	y1, m1, _ := t.Date()
	t = t.AddDate(0, month, 0)
	y2, m2, _ := t.Date()
	day := 1
	if (y2-y1)*12+int(m2-m1) > month {
		day = t.Day()
	}
	return t.AddDate(0, 0, -day)
}

// AddMonth 获取几个月后的日期
func AddMonth(t int64, month int) int64 {
	return AddMonthTime(time.Unix(t, 0), month).Unix()
}

func TestAddMonth() {
	st := time.Date(2018, 1, 28, 0, 0, 0, 0, time.Local)
	am := AddMonthTime(st, 1)
	fmt.Println("One month after ", st, " is ", am)
	fmt.Println("-------------------------------")

	st = time.Date(2018, 1, 29, 0, 0, 0, 0, time.Local)
	am = AddMonthTime(st, 1)
	fmt.Println("One month after ", st, " is ", am)
	fmt.Println("-------------------------------")

	st = time.Date(2018, 1, 30, 0, 0, 0, 0, time.Local)
	am = AddMonthTime(st, 1)
	fmt.Println("One month after ", st, " is ", am)
	fmt.Println("-------------------------------")

	st = time.Date(2018, 1, 31, 0, 0, 0, 0, time.Local)
	am = AddMonthTime(st, 1)
	fmt.Println("One month after ", st, " is ", am)
	fmt.Println("-------------------------------")

	st = time.Date(2018, 2, 1, 0, 0, 0, 0, time.Local)
	am = AddMonthTime(st, 1)
	fmt.Println("One month after ", st, " is ", am)
	fmt.Println("-------------------------------")

	st = time.Date(2020, 1, 28, 0, 0, 0, 0, time.Local)
	am = AddMonthTime(st, 1)
	fmt.Println("One month after ", st, " is ", am)
	fmt.Println("-------------------------------")

	st = time.Date(2020, 1, 29, 0, 0, 0, 0, time.Local)
	am = AddMonthTime(st, 1)
	fmt.Println("One month after ", st, " is ", am)
	fmt.Println("-------------------------------")

	st = time.Date(2020, 1, 30, 0, 0, 0, 0, time.Local)
	am = AddMonthTime(st, 1)
	fmt.Println("One month after ", st, " is ", am)
	fmt.Println("-------------------------------")

	st = time.Date(2020, 1, 31, 0, 0, 0, 0, time.Local)
	am = AddMonthTime(st, 1)
	fmt.Println("One month after ", st, " is ", am)
	fmt.Println("-------------------------------")

	st = time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local)
	am = AddMonthTime(st, 1)
	fmt.Println("One month after ", st, " is ", am)
	fmt.Println("-------------------------------")

}
