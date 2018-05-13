package main

import (
	"fmt"
	"time"
)

/*
var transMonth = map[time.Month]int{
	time.January:   1,
	time.February:  2,
	time.March:     3,
	time.April:     4,
	time.May:       5,
	time.June:      6,
	time.July:      7,
	time.August:    8,
	time.September: 9,
	time.October:   10,
	time.November:  11,
	time.December:  12,
}

// AddMonthTime222 获取几个月后的日期
func AddMonthTime222(t time.Time, month int) time.Time {
	y1, m1, _ := t.Date()
	if m1 == 2 {
		if month > 1 {
			t = t.AddDate(0, 1, 0)
			if t == GetEndDateOfCurrentMonth2(t.Unix()) {
				t = GetStartDateOfNextMonth(t.Unix())
			}
			t = t.AddDate(0, month-1, 0)
		} else {
			t = t.AddDate(0, month, 0)
		}
	} else {
		t = t.AddDate(0, month, 0)
	}

	y2, m2, _ := t.Date()
	day := 1
	if (y2-y1)*12+int(m2-m1) > month {
		day = t.Day()
	}
	return t.AddDate(0, 0, -day)
}

// Date returns the year, month, and day in which t occurs.
func InitDate(t time.Time) (year, month, day int) {
	t1 := t.AddDate(0, month, 0)
	year, m, _ := t1.Date()
	month = transMonth(m)

	return year, month, day
}
*/

//2017-06-30 2018-06-29 => 4 2017-12-30 2018-03-29

// AddMonthTime 获取几个月后的日期(加付几减一天,如果不存在这一天,推到最近一天)
func AddMonthTime22(t time.Time, month int) time.Time {
	//t1 := t.AddDate(0, month, 0)
	t1 := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local).AddDate(0, month, 0)
	fmt.Println("t1:", t1)
	y, m, _ := t1.Date()
	_, _, d := t.Date()
	//fmt.Println("y,m,d:", y, m, d)

	t2 := time.Date(y, m, d, 0, 0, 0, 0, time.Local)

	_, m2, _ := t2.Date()
	//fmt.Println("m2:", m2)
	if m != m2 {
		return GetEndDateOfCurrentMonth2(t1.Unix())
	}
	return time.Date(y, m, d-1, 0, 0, 0, 0, time.Local)
}

/*
// MakeUpDaysTime2 补足天数后账期的起始/终止日期
func MakeUpDaysTime22(startTime int64, addMonths, addDays, paymentMonth int) (time.Time, time.Time) {
	dayAfterMakeUp, months := MakeUpCurrentMonth22(startTime, addMonths, addDays, paymentMonth)

	// 改造
	//	return time.Unix(startTime, 0), AddMonthTime22(dayAfterMakeUp, months)

}


// MakeUpDays22 补足天数后账期的起始/终止日期
func MakeUpDays22(startTime int64, addMonths, addDays, paymentMonth int) (int64, int64) {
	start, end := MakeUpDaysTime22(startTime, addMonths, addDays, paymentMonth)
	return start.Unix(), end.Unix()
}

// MakeUpCurrentMonth 补足当月
func MakeUpCurrentMonth22(startTime int64, addMonths, addDays, paymentMonth int) (dayAfterMakeUp time.Time, months int) {
	year, month, _ := time.Unix(startTime, 0).Date()

	if addDays == 0 {
		dayAfterMakeUp = time.Unix(startTime, 0)
		if addMonths == 0 {
			months = paymentMonth
		} else {
			months = addMonths + paymentMonth - 1
		}
	} else {
		dayAfterMakeUp = time.Date(year, month, addDays, 0, 0, 0, 0, time.Local)
		months = addMonths + paymentMonth - 1
	}
	if addDays == 1 {
		dayAfterMakeUp = dayAfterMakeUp.AddDate(0, 1, 0)
	}

	return dayAfterMakeUp, months
}

// RecursiveCalcPeriod 循环计算期账单账期
func RecursiveCalcPeriod22(startTime int64, addMonths, addDays, paymentMonth, term int) (periodStart, periodEnd int64) {
	tmpStart, tmpEnd := MakeUpDays2(startTime, addMonths, addDays, paymentMonth)
	if term == 1 {
		return tmpStart, tmpEnd
	}

	for i := 2; i <= term; i++ {
		periodStart, periodEnd = CalcNextBaseOnPre22(tmpStart, tmpEnd, paymentMonth)
		tmpStart = periodStart
		tmpEnd = periodEnd
	}

	return periodStart, periodEnd
}

// CalcNextBaseOnPre22 根据上一个账期推算出下一个账期
func CalcNextBaseOnPre22(startTime, endTime int64, paymentMonth int) (periodStart, periodEnd int64) {
	periodStart = GetLaterDateTime(time.Unix(endTime, 0), 1)
	periodEnd = AddMonth(periodStart, paymentMonth)
	return periodStart, periodEnd

	t := time.Unix(startTime, 0)
	//	t := t.AddDate(0, 1, 0)
	return
}
*/
