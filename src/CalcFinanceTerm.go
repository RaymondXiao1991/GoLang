package main

import (
	"fmt"
	"time"
)

// CalcFinanceFirstTerm 计算金融产品首账单
func CalcFinanceFirstTerm(startTime int64, firstMonth, firstDay, paymentMonth int) time.Time {
	addMonth := 0
	addDay := 0
	if firstDay == 0 {
		addMonth = firstMonth + paymentMonth - 1
	} else if firstDay < time.Unix(startTime, 0).Day() {
		addMonth = firstMonth + paymentMonth - 1
		addDay = -firstDay
	} else {
		addMonth = firstMonth + paymentMonth
	}

	return AddMonthTime(time.Unix(startTime, 0), addMonth).AddDate(0, 0, addDay)
}

func CalcFinanceFirstTermStr(start time.Time, firstMonth, firstDay, paymentMonth int) time.Time {
	e := CalcFinanceFirstTerm2(start.Unix(), firstMonth, firstDay, paymentMonth)
	return e
}

func CalcFinancePeriodTermStr(start time.Time, firstMonth, firstDay, paymentMonth, term int) (time.Time, time.Time) {
	s, e := CalcFinancePeriodTerm(start.Unix(), firstMonth, firstDay, paymentMonth, term)
	return s, e
}

func CalcFinanceLastTermStr(start time.Time, firstMonth, firstDay, paymentMonth, totalTerm int) (time.Time, time.Time) {
	s, e := CalcFinanceLastTerm(start.Unix(), firstMonth, firstDay, paymentMonth, totalTerm)
	return s, e
}

func TestCase4(start string, firstMonth, firstDay, paymentMonth int) {
	t1, _ := time.Parse("2006-01-02", start)
	e := CalcFinanceFirstTermStr(t1, firstMonth, firstDay, paymentMonth)
	fmt.Println(start, "=>", e.Format("2006-01-02"))
}

func TestCase5(start string, firstMonth, firstDay, paymentMonth, term int) {
	t1, _ := time.Parse("2006-01-02", start)
	s, e := CalcFinancePeriodTermStr(t1, firstMonth, firstDay, paymentMonth, term)
	fmt.Println(start, "=>", s.Format("2006-01-02"), "~", e.Format("2006-01-02"))
}

func TestCase6(start string, firstMonth, firstDay, paymentMonth, totalTerm int) {
	t1, _ := time.Parse("2006-01-02", start)
	s, e := CalcFinanceLastTermStr(t1, firstMonth, firstDay, paymentMonth, totalTerm)
	fmt.Println(start, "=>", s.Format("2006-01-02"), "~", e.Format("2006-01-02"))
}

// CalcFinanceFirstTerm2 计算金融产品首账单
func CalcFinanceFirstTerm2(startTime int64, firstMonth, firstDay, paymentMonth int) time.Time {
	var termEndTime time.Time
	if firstDay >= 1 {
		startDateOfNextMonth := GetStartDateOfNextMonth(startTime)
		startDateOfNextMonth = AddMonthTime(time.Unix(startDateOfNextMonth.Unix(), 0), paymentMonth+firstMonth-2)
		//fmt.Println("startDateOfNextMonth:", startDateOfNextMonth.Format("2006-01-02"))
		termEndTime = GetLaterDate(startDateOfNextMonth, firstDay-1)
	} else {
		termEndTime = AddMonthTime(time.Unix(startTime, 0), paymentMonth+firstMonth-1)
		//fmt.Println("termEndTime:", termEndTime.Format("2006-01-02"))
	}

	return termEndTime
}

// CalcFinancePeriodTerm 计算金融产品期账单
func CalcFinancePeriodTerm(startTime int64, firstMonth, firstDay, paymentMonth, term int) (time.Time, time.Time) {
	var (
		termStartTime, termEndTime time.Time
	)
	if firstDay >= 1 {
		startDateOfNextMonth := GetStartDateOfNextMonth(startTime)
		startDateOfNextMonth = AddMonthTime(time.Unix(startDateOfNextMonth.Unix(), 0), paymentMonth*term+firstMonth-2)
		termStartTime = GetLaterDate(startDateOfNextMonth, firstDay)
		termEndTime = AddMonthTime(termStartTime, paymentMonth*term)
	} else {
		termStartTime = AddMonthTime(GetStartDateOfNextMonth(startTime), paymentMonth*(term-1)).AddDate(0, 0, 1)
		termEndTime = AddMonthTime(GetStartDateOfNextMonth(startTime), paymentMonth*term)
	}

	return termStartTime, termEndTime
}

// CalcStartEndTimeOfIrrTerm 计算租期起始时间(不规则账期)
func CalcStartEndTimeOfIrrTermXxx(startTime, endTime int64, term, totalTerm, billType, paymentMonth int, isTermNormal bool) (termStartTime, termEndTime int64) {
	endDateOfStartTime := GetStartDateOfNextMonth(startTime).Unix()
	fmt.Println("endDateOfStartTime:", endDateOfStartTime)
	switch billType {
	case 1:
		termStartTime = startTime
		if isTermNormal {
			termEndTime = AddMonth(endDateOfStartTime, paymentMonth)
		} else {
			termEndTime = endTime
		}
		fmt.Println("termEndTime:", termEndTime)
	case 2:
		termStartTime = AddMonthTime(time.Unix(endDateOfStartTime, 0), paymentMonth*(term-1)).AddDate(0, 0, 1).Unix()
		if isTermNormal {
			termEndTime = AddMonth(endDateOfStartTime, paymentMonth*term)
		} else {
			termEndTime = endTime
		}
		fmt.Println("termStartTime:", termStartTime)
		fmt.Println("termEndTime:", termEndTime)
	case 3:
		if isTermNormal {
			termStartTime = AddMonth(endDateOfStartTime, paymentMonth*(totalTerm-1))
			termEndTime = AddMonth(endDateOfStartTime, paymentMonth*(totalTerm))
		} else {
			termStartTime = CurrentTimeStamp()
			termEndTime = CurrentTimeStamp()
		}
		fmt.Println("termStartTime:", termStartTime)
		fmt.Println("termEndTime:", termEndTime)
	case 4:
		termStartTime = CurrentTimeStamp()
		termEndTime = CurrentTimeStamp()
	case 5:
		termStartTime = CurrentTimeStamp()
		termEndTime = CurrentTimeStamp()
	}

	return
}

// CalcFinanceLastTerm 计算金融产品末账单
func CalcFinanceLastTerm(startTime int64, firstMonth, firstDay, paymentMonth, totalTerm int) (time.Time, time.Time) {
	var (
		termStartTime, termEndTime time.Time
	)
	if firstDay >= 1 {
		startDateOfNextMonth := GetStartDateOfNextMonth(startTime)
		tempStartTime := AddMonthTime(time.Unix(startDateOfNextMonth.Unix(), 0), paymentMonth*(totalTerm-2)+firstMonth-2)
		termStartTime = GetLaterDate(tempStartTime, firstDay)
		termEndTime = AddMonthTime(termStartTime, paymentMonth)
	} else {
		termStartTime = AddMonthTime(time.Unix(startTime, 0), paymentMonth*totalTerm+firstMonth)
		termEndTime = termStartTime
	}

	return termStartTime, termEndTime
}