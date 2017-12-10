package main  

import (  
	"fmt"  
	"time"  
)  

// 求两个时间之间的时间差几个月零几天
func CalcTimeBetweenDates(startDate int64, endDate int64) (months, days int) {
	s := time.Unix(startDate, 0)
	e := time.Unix(endDate, 0)

	m := int(endDate - startDate)/(24*60*60)/31
	fmt.Println("m:", m)
	d := int(endDate - (s.AddDate(0,int(m),0)).Unix())/(24*60*60)
	fmt.Println("d:", d)

	ds := DayInMonth(e)

	fmt.Println("ds:", ds)

	if d > ds {
		m++
		d = d - ds
	}

	return m, d
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
	fmt.Println(MonthDays[t.Month()])
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

