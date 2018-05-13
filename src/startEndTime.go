package main

import (
	"fmt"
	"sync"
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

// MakeUpDaysTime 补足天数
func MakeUpDaysTime(now int64, addMonths, addDays, paymentMonth int) time.Time {
	year, month, _ := time.Unix(now, 0).Date()
	if addDays == 1 {
		return AddMonthTime(time.Date(year, month, 1, 0, 0, 0, 0, time.Local).AddDate(0, 1, 0), addMonths+paymentMonth-1)
	} else if addDays > 1 {
		return AddMonthTime(time.Date(year, month, addDays, 0, 0, 0, 0, time.Local), addMonths+paymentMonth-1)
	} else {
		if addMonths == 0 {
			return AddMonthTime(time.Unix(now, 0), paymentMonth)
		}
		return AddMonthTime(time.Unix(now, 0), addMonths+paymentMonth-1)
	}
}

// MakeUpDays 补足天数
func MakeUpDays(startTime int64, addMonths, addDays, paymentMonth int) int64 {
	return MakeUpDaysTime(startTime, addMonths, addDays, paymentMonth).Unix()
}

// MakeUpDaysTime2 补足天数后账期的起始/终止日期
func MakeUpDaysTime2(startTime int64, addMonths, addDays, paymentMonth int) (time.Time, time.Time) {
	dayAfterMakeUp, months := MakeUpCurrentMonth(startTime, addMonths, addDays, paymentMonth)
	return time.Unix(startTime, 0), AddMonthTime(dayAfterMakeUp, months)
}

// MakeUpDays2 补足天数后账期的起始/终止日期
func MakeUpDays2(startTime int64, addMonths, addDays, paymentMonth int) (int64, int64) {
	start, end := MakeUpDaysTime2(startTime, addMonths, addDays, paymentMonth)
	return start.Unix(), end.Unix()
}

// MakeUpCurrentMonth 补足当月
func MakeUpCurrentMonth(startTime int64, addMonths, addDays, paymentMonth int) (dayAfterMakeUp time.Time, months int) {
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

// CalcPeriod 期账单账期
func CalcPeriod(startTime int64, addMonths, addDays, paymentMonth, term int) (periodStart, periodEnd int64) {
	//year, month, _ := MakeUpDaysTime(startTime, addMonths, addDays, paymentMonth).Date()
	//periodStart = time.Date(year, month, 1, 0, 0, 0, 0, time.Local).Unix()
	//endOfFirstPeriod := MakeUpDaysTime(startTime, addMonths, addDays, paymentMonth)
	//thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	//temp := thisMonth.AddDate(0, 1, 0)

	//year, month, day := MakeUpDaysTime(startTime, addMonths, addDays, paymentMonth).Date()
	//endOfFirstPeriod := time.Date(year, month, day+1, 0, 0, 0, 0, time.Local)
	endOfFirstPeriod := GetLaterDate(MakeUpDaysTime(startTime, addMonths, addDays, paymentMonth), 1)
	//periodStart = AddMonthTime(endOfFirstPeriod, paymentMonth*(term-2)).AddDate(0, 0, 1).Unix()
	//fmt.Println("endOfFirstPeriod:", endOfFirstPeriod)
	periodStart = GetLaterDateTime(AddMonthTime(endOfFirstPeriod, paymentMonth*(term-2)), 1)
	periodEnd = AddMonth(periodStart, paymentMonth)
	return periodStart, periodEnd
}

// RecursiveCalcPeriod2 递归计算期账单账期
func RecursiveCalcPeriod2(startTime int64, addMonths, addDays, paymentMonth, term int) (periodStart, periodEnd int64) {
	if term == 1 {
		periodStart = startTime
		periodEnd = MakeUpDays(startTime, addMonths, addDays, paymentMonth)
		return periodStart, periodEnd
	}

	_, lastEndTime := RecursiveCalcPeriod2(startTime, addMonths, addDays, paymentMonth, term-1)
	periodStart = GetLaterDateTime(time.Unix(lastEndTime, 0), 1)
	periodEnd = AddMonth(periodStart, paymentMonth)
	return periodStart, periodEnd
}

// RecursiveCalcPeriod3 递归计算期账单账期
func RecursiveCalcPeriod3(startTime int64, addMonths, addDays, paymentMonth, term int) (periodStart, periodEnd int64) {
	var tmpEnd int64

	tmpEnd = MakeUpDays(startTime, addMonths, addDays, paymentMonth)
	if term == 1 {
		return startTime, tmpEnd
	} else {
		periodStart = AddMonth(GetLaterDateTime(time.Unix(tmpEnd, 0), term-1), paymentMonth*(term-2))
		periodEnd = AddMonth(periodStart, paymentMonth)
		return periodStart, periodEnd
	}
}

// RecursiveCalcPeriod4 递归计算期账单账期
func RecursiveCalcPeriod4(startTime int64, addMonths, addDays, paymentMonth, term int) (periodStart, periodEnd int64) {
	if term == 1 {
		return startTime, MakeUpDays(startTime, addMonths, addDays, paymentMonth)
	}

	var w sync.WaitGroup
	var lastEndTime int64
	w.Add(1)
	go func() {
		defer w.Done()
		_, lastEndTime = RecursiveCalcPeriod4(startTime, addMonths, addDays, paymentMonth, term-1)
	}()
	w.Wait()

	periodStart = GetLaterDateTime(time.Unix(lastEndTime, 0), 1)
	periodEnd = AddMonth(periodStart, paymentMonth)
	return periodStart, periodEnd
}

// RecursiveCalcPeriod 循环计算期账单账期
func RecursiveCalcPeriod(startTime int64, addMonths, addDays, paymentMonth, term int) (periodStart, periodEnd int64) {
	tmpStart, tmpEnd := MakeUpDays2(startTime, addMonths, addDays, paymentMonth)
	if term == 1 {
		return tmpStart, tmpEnd
	}

	for i := 2; i <= term; i++ {
		periodStart, periodEnd = CalcNextBaseOnPre(tmpStart, tmpEnd, paymentMonth*i)
		tmpStart = periodStart
		tmpEnd = periodEnd
	}

	return periodStart, periodEnd
}

// CalcNextBaseOnPre 根据上一个账期推算出下一个账期
func CalcNextBaseOnPre(startTime, endTime int64, paymentMonth int) (periodStart, periodEnd int64) {
	periodStart = GetLaterDateTime(time.Unix(endTime, 0), 1)
	//periodEnd = AddMonth(periodStart, paymentMonth)
	periodEnd = AddMonth(periodStart, paymentMonth)
	return periodStart, periodEnd
}

// OptimizedRecursiveCalcPeriod 递归计算期账单账期(优化)
func OptimizedRecursiveCalcPeriod(startTime int64, addMonths, addDays, paymentMonth, term int) (periodStart, periodEnd int64) {
	if term == 1 {
		periodStart = startTime
		periodEnd = MakeUpDays(startTime, addMonths, addDays, paymentMonth)
		return periodStart, periodEnd
	}

	_, lastEndTime := OptimizedRecursiveCalcPeriod(startTime, addMonths, addDays, paymentMonth, term-1)
	periodStart = GetLaterDateTime(time.Unix(lastEndTime, 0), 1)
	periodEnd = AddMonth(periodStart, paymentMonth)
	return periodStart, periodEnd
	/*
	   	int* f=new int[n+1];
	       f[1]=1;f[0]=0;
	       for(int i=0;i<=n;i++);
	           f[i]=f[i-1]+f[i-2];
	       int r=f[n];
	       delete[] f;
	   	return r;
	*/
}

// CalcLast 结算账单账期
func CalcLast(startTime int64, addMonths, addDays, paymentMonth, totalTerm int) (periodStart, periodEnd int64) {
	return CalcPeriod(startTime, addMonths, addDays, paymentMonth, totalTerm-1)
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

/*
// AddMonthTime33 获取几个月后的日期(加付几减一天,如果不存在这一天,推到最近一天)
func AddMonthTime33(t1, t2 time.Time, months int) time.Time {
	//t1 := t.AddDate(0, month, 0)
	t1 := time.Date(t1.Year(), t1.Month(), 1, 0, 0, 0, 0, time.Local).AddDate(0, months, 0)
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

// AddMonth 获取几个月后的日期
func AddMonth33(t1, t2 int64, month int) int64 {
	return AddMonthTime33(time.Unix(t1, 0), time.Unix(t2, 0), month).Unix()
}
*/
