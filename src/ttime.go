package main

import (
	"fmt"
	"time"
)

// 求两个时间之间的时间差几个月零几天
func CalcTimeBetweenDates(startDate int64, endDate int64) (months, days int) {
	start := time.Unix(startDate, 0)
	months = int(endDate-startDate) / (24 * 60 * 60) / 31
	later := start.AddDate(0, int(months), 0)
	days = int(endDate-later.Unix())/(24*60*60) + 1
	dayInMonth := DayInMonth(later)
	if days >= dayInMonth {
		months++
		days = days - dayInMonth
	}
	return months, days
}

func IsLeapYear(year int) bool {
	isLeapYear := year%4 == 0 && (year%100 != 0 || year%400 == 0)
	// 闰年366
	if isLeapYear {
		return true
	}
	// 平年365天
	return false
}

func DayInMonth(t time.Time) int {
	if t.Month() == 2 {
		if IsLeapYear(t.Year()) {
			return 29
		} else {
			return 28
		}
	}
	return MonthDays[t.Month()]
}

var MonthDays = map[time.Month]int{
	1:  31,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

func TestCaseOfCalcTime() {
	// 1515254400 2018-01-07 00:00:00
	// 1517932800 2018-02-07 00:00:00
	fmt.Println("2018-01-07 00:00:00 -- 2018-02-07 00:00:00")
	m, d := CalcTimeBetweenDates(1515254400, 1517932800)
	fmt.Println("m,d:", m, d)
	fmt.Println("--------------------")

	// 1485960657 2017-2-1 22:50:57
	// 1512917457 2017-12-10 22:50:57
	fmt.Println("2017-2-1 22:50:57 -- 2017-12-10 22:50:57")
	m, d = CalcTimeBetweenDates(1485960657, 1512917457)
	fmt.Println("m,d:", m, d)
	fmt.Println("--------------------")

	// 1485960657 2017-2-1 22:50:57
	// 1514645457 2017-12-30 22:50:57
	fmt.Println("2017-2-1 22:50:57 -- 2017-12-30 22:50:57")
	m, d = CalcTimeBetweenDates(1485960657, 1514645457)
	fmt.Println("m,d:", m, d)
	fmt.Println("--------------------")

	// 1449800745 2015/12/11 10:25:45
	// 1514645457 2017-12-30 22:50:57
	fmt.Println("2015/12/11 10:25:45 -- 2017-12-30 22:50:57")
	m, d = CalcTimeBetweenDates(1449800745, 1514645457)
	fmt.Println("m,d:", m, d)
	fmt.Println("--------------------")

	// 1483200000 2017/1/1 00:00:00
	// 1488384000 2017/3/2 00:00:00
	fmt.Println("2017/1/1 00:00:00 -- 2017/3/2 00:00:00")
	m, d = CalcTimeBetweenDates(1483200000, 1488384000)
	fmt.Println("m,d:", m, d)
	fmt.Println("--------------------")

	// 1485878400 2017/2/1 00:00:00
	// 1490976000 2017/4/1 00:00:00
	fmt.Println("2017/2/1 00:00:00 -- 2017/4/1 00:00:00")
	m, d = CalcTimeBetweenDates(1485878400, 1490976000)
	fmt.Println("m,d:", m, d)
	fmt.Println("--------------------")

	// 1485964800 2017/2/2 00:00:00
	// 1490976000 2017/4/1 00:00:00
	fmt.Println("2017/2/2 00:00:00 -- 2017/4/1 00:00:00")
	m, d = CalcTimeBetweenDates(1485964800, 1490976000)
	fmt.Println("m,d:", m, d)
	fmt.Println("--------------------")

	// 1072886400 2004/1/1 00:00:00
	// 1078070400 2004/3/1 00:00:00
	fmt.Println("2004/1/1 00:00:00 -- 2004/3/1 00:00:00")
	m, d = CalcTimeBetweenDates(1072886400, 1078070400)
	fmt.Println("m,d:", m, d)
	fmt.Println("--------------------")

	// 1072886400 2004/1/1 00:00:00
	// 1077984000 2004/2/29 00:00:00
	fmt.Println("2004/1/1 00:00:00 -- 2004/2/29 00:00:00")
	m, d = CalcTimeBetweenDates(1072886400, 1077984000)
	fmt.Println("m,d:", m, d)
	fmt.Println("--------------------")

	// 1513094400 2017/12/13 0:0:0
	// 1519660800 2018/2/27 0:0:0
	fmt.Println("2017/12/13 0:0:0 -- 2018/2/27 0:0:0")
	m, d = CalcTimeBetweenDates(1513094400, 1519660800)
	fmt.Println("m,d:", m, d)
	fmt.Println("--------------------")

	// 1503072000 2017/08/19 0:0:0
	// 1514044800 2017/12/24 0:0:0
	fmt.Println("2017/08/19 0:0:0 -- 2017/12/24 0:0:0")
	m, d = CalcTimeBetweenDates(1503072000, 1514044800)
	fmt.Println("m,d:", m, d)
	fmt.Println("--------------------")
}
