package main

import (
	"fmt"
	"time"
)

func CurrentTimeStamp() int64 {
	return time.Now().Unix()
}

// CalTermEndTime 计算租期结束时间
func CalTermEndTime(startTime, endTime int64, term, totalTerm, billType, paymentMonth int, isTermNormal bool) (startDateStr, endDateStr int64) {
	switch billType {
	case 1:
		startDateStr = startTime
		if isTermNormal {
			endDateStr = AddMonth(startTime, paymentMonth)
		} else {
			endDateStr = endTime
		}
	case 2:
		startDateStr = AddMonthTime(time.Unix(startTime, 0), paymentMonth*(term-1)).AddDate(0, 0, 1).Unix()
		if isTermNormal {
			endDateStr = AddMonth(startTime, paymentMonth*term)
		} else {
			endDateStr = endTime
		}
	case 3:
		if isTermNormal {
			startDateStr = AddMonth(startTime, paymentMonth*(totalTerm-1))
			endDateStr = AddMonth(startTime, paymentMonth*(totalTerm))
		} else {
			startDateStr = CurrentTimeStamp()
			endDateStr = CurrentTimeStamp()
		}
	case 4:
		startDateStr = CurrentTimeStamp()
		endDateStr = CurrentTimeStamp()
	case 5:
		startDateStr = CurrentTimeStamp()
		endDateStr = CurrentTimeStamp()
	}

	return
}

// CalcStartEndTimeOfIrrTerm 计算租期起始时间(不规则账期)
func CalcStartEndTimeOfIrrTerm(startTime, endTime int64, term, totalTerm, billType, paymentMonth int, isTermNormal bool) (termStartTime, termEndTime int64) {
	endDateOfStartTime := GetStartDateOfNextMonth(startTime).Unix()
	switch billType {
	case 1:
		termStartTime = startTime
		if isTermNormal {
			termEndTime = AddMonth(endDateOfStartTime, paymentMonth)
		} else {
			termEndTime = endTime
		}
	case 2:
		termStartTime = AddMonthTime(time.Unix(endDateOfStartTime, 0), paymentMonth*(term-1)).AddDate(0, 0, 1).Unix()
		if isTermNormal {
			termEndTime = AddMonth(endDateOfStartTime, paymentMonth*term)
		} else {
			termEndTime = endTime
		}
	case 3:
		if isTermNormal {
			termStartTime = AddMonth(endDateOfStartTime, paymentMonth*(totalTerm-1))
			termEndTime = AddMonth(endDateOfStartTime, paymentMonth*(totalTerm))
		} else {
			termStartTime = CurrentTimeStamp()
			termEndTime = CurrentTimeStamp()
		}
	case 4:
		termStartTime = CurrentTimeStamp()
		termEndTime = CurrentTimeStamp()
	case 5:
		termStartTime = CurrentTimeStamp()
		termEndTime = CurrentTimeStamp()
	}

	return
}

// CalcStartEndTime 计算租期
func CalcStartEndTime(startTime, endTime int64, term, totalTerm, billType, paymentMonth int, firstMonth, firstDay int, isTermNormal bool) (termStartTime, termEndTime int64) {
	switch billType {
	case 1:
		termStartTime = startTime
		if isTermNormal {
			termEndTime = MakeUpDays(startTime, firstMonth, firstDay, paymentMonth)
		} else {
			termEndTime = endTime
		}
	case 2:
		//termStartTime, termEndTime = CalcPeriod(startTime, firstMonth, firstDay, paymentMonth, term)
		termStartTime, termEndTime = RecursiveCalcPeriod(startTime, firstMonth, firstDay, paymentMonth, term)
		if !isTermNormal {
			termEndTime = endTime
		}
	case 3:
		termStartTime = CurrentTimeStamp()
		termEndTime = CurrentTimeStamp()
	case 4:
		termStartTime = CurrentTimeStamp()
		termEndTime = CurrentTimeStamp()
	case 5:
		termStartTime = CurrentTimeStamp()
		termEndTime = CurrentTimeStamp()
	}

	return
}

func CalcStartEndTimeStr(startTime, endTime time.Time, term, totalTerm, billType, paymentMonth, firstMonth, firstDay int, isTermNormal bool) (startDateStr, endDateStr time.Time) {
	s, e := CalcStartEndTime(startTime.Unix(), endTime.Unix(), term, totalTerm, billType, paymentMonth, firstMonth, firstDay, isTermNormal)
	return time.Unix(s, 0), time.Unix(e, 0)
}

func TestCalcStartEndTimeStr(start, end string, term, totalTerm, billType, paymentMonth, firstMonth, firstDay int, isTermNormal bool) {
	t1, _ := time.Parse("2006-01-02", start)
	t2, _ := time.Parse("2006-01-02", end)
	s, e := CalcStartEndTimeStr(t1, t2, term, totalTerm, billType, paymentMonth, firstMonth, firstDay, isTermNormal)
	fmt.Println(start, end, "=>", term, s.Format("2006-01-02"), e.Format("2006-01-02"))
}

func CalTermEndTimeStr(startTime, endTime time.Time, term, totalTerm, billType, paymentMonth int, isTermNormal bool) (startDateStr, endDateStr time.Time) {
	s, e := CalTermEndTime(startTime.Unix(), endTime.Unix(), term, totalTerm, billType, paymentMonth, isTermNormal)
	return time.Unix(s, 0), time.Unix(e, 0)
}

func CalcStartEndTimeOfIrrTermStr(startTime, endTime time.Time, term, totalTerm, billType, paymentMonth int, isTermNormal bool) (startDateStr, endDateStr time.Time) {
	s, e := CalcStartEndTimeOfIrrTerm(startTime.Unix(), endTime.Unix(), term, totalTerm, billType, paymentMonth, isTermNormal)
	return time.Unix(s, 0), time.Unix(e, 0)
}

func TestCase(start, end string, term, totalTerm, billType, paymentMonth int, isTermNormal bool) {
	t1, _ := time.Parse("2006-01-02", start)
	t2, _ := time.Parse("2006-01-02", end)
	s, e := CalTermEndTimeStr(t1, t2, term, totalTerm, billType, paymentMonth, isTermNormal)
	fmt.Println(start, end, "=>", term, s.Format("2006-01-02"), e.Format("2006-01-02"))
}

func CalTermEndTimeStr2(startTime, endTime time.Time, term, totalTerm, billType, paymentMonth int, isTermNormal bool) (startDateStr, endDateStr time.Time) {
	s, e := CalcStartEndTimeOfIrrTerm(startTime.Unix(), endTime.Unix(), term, totalTerm, billType, paymentMonth, isTermNormal)
	return time.Unix(s, 0), time.Unix(e, 0)
}

func TestCase2(start, end string, term, totalTerm, billType, paymentMonth int, isTermNormal bool) {
	t1, _ := time.Parse("2006-01-02", start)
	t2, _ := time.Parse("2006-01-02", end)
	s, e := CalcStartEndTimeOfIrrTermStr(t1, t2, term, totalTerm, billType, paymentMonth, isTermNormal)
	fmt.Println(start, end, "=>", term, s.Format("2006-01-02"), e.Format("2006-01-02"))
}

func TestCaseOfMakeUpDays(start string, addMonths, addDays, paymentMonth int) {
	t1, _ := time.Parse("2006-01-02", start)
	e := MakeUpDaysTime(t1.Unix(), addMonths, addDays, paymentMonth)
	fmt.Println(start, addMonths, addDays, paymentMonth, "=>", start, e.Format("2006-01-02"))
}

func TestCaseOfMakeUpDaysStr() {
	TestCaseOfMakeUpDays("2017-04-03", 1, 1, 1)
	TestCaseOfMakeUpDays("2017-04-03", 1, 18, 1)
	TestCaseOfMakeUpDays("2017-04-03", 1, 0, 1)
	TestCaseOfMakeUpDays("2017-04-03", 0, 0, 1)
	TestCaseOfMakeUpDays("2017-04-03", 0, 0, 2)

	TestCaseOfMakeUpDays("2017-04-01", 1, 0, 1)
	TestCaseOfMakeUpDays("2017-04-02", 1, 0, 1)
	TestCaseOfMakeUpDays("2017-02-27", 1, 0, 1)
	TestCaseOfMakeUpDays("2017-02-28", 1, 0, 1)
	TestCaseOfMakeUpDays("2017-04-01", 1, 1, 1)
	TestCaseOfMakeUpDays("2017-04-02", 1, 1, 1)
	TestCaseOfMakeUpDays("2017-02-27", 1, 1, 1)
	TestCaseOfMakeUpDays("2017-02-28", 1, 1, 1)
	TestCaseOfMakeUpDays("2017-04-01", 0, 0, 1)
	TestCaseOfMakeUpDays("2017-04-02", 0, 0, 1)
	TestCaseOfMakeUpDays("2017-02-27", 0, 0, 1)
	TestCaseOfMakeUpDays("2017-02-28", 0, 0, 1)
	TestCaseOfMakeUpDays("2017-04-01", 1, 18, 1)
	TestCaseOfMakeUpDays("2017-04-02", 1, 18, 1)
	TestCaseOfMakeUpDays("2017-02-27", 1, 18, 1)
	TestCaseOfMakeUpDays("2017-02-28", 1, 18, 1)
}

func TestCalTermEndTimeStr() {

	TestCase("2017-04-01", "2017-05-31", 1, 2, 1, 1, true)
	TestCase("2017-04-01", "2017-05-31", 2, 2, 2, 1, true)

	TestCase("2017-01-01", "2017-02-31", 1, 2, 1, 1, true)
	TestCase("2017-01-01", "2017-02-31", 2, 2, 2, 1, true)

	TestCase("2017-01-02", "2017-09-01", 1, 8, 1, 1, true)
	TestCase("2017-01-02", "2017-09-01", 2, 8, 2, 1, true)
	TestCase("2017-01-02", "2017-09-01", 3, 8, 2, 1, true)
	TestCase("2017-01-02", "2017-09-01", 4, 8, 2, 1, true)
	TestCase("2017-01-02", "2017-09-01", 5, 8, 2, 1, true)
	TestCase("2017-01-02", "2017-09-01", 6, 8, 2, 1, true)
	TestCase("2017-01-02", "2017-09-01", 7, 8, 2, 1, true)
	TestCase("2017-01-02", "2017-09-01", 8, 8, 2, 1, true)
	TestCase("2017-01-02", "2017-09-01", 8, 8, 3, 1, false)

	TestCase("2017-11-26", "2018-01-25", 1, 2, 1, 1, true)
	TestCase("2017-11-26", "2018-01-25", 2, 2, 2, 1, true)

	TestCase("2017-11-26", "2018-01-26", 1, 3, 1, 1, true)
	TestCase("2017-11-26", "2018-01-26", 2, 3, 2, 1, true)
	TestCase("2017-11-26", "2018-01-26", 3, 3, 2, 1, false)

	TestCase("2017-03-01", "2017-10-31", 1, 8, 1, 1, true)
	TestCase("2017-03-01", "2017-10-31", 2, 8, 2, 1, true)
	TestCase("2017-03-01", "2017-10-31", 3, 8, 2, 1, true)
	TestCase("2017-03-01", "2017-10-31", 4, 8, 2, 1, true)
	TestCase("2017-03-01", "2017-10-31", 5, 8, 2, 1, true)
	TestCase("2017-03-01", "2017-10-31", 6, 8, 2, 1, true)
	TestCase("2017-03-01", "2017-10-31", 7, 8, 2, 1, true)
	TestCase("2017-03-01", "2017-10-31", 8, 8, 2, 1, true)

	TestCase("2017-02-02", "2018-03-01", 1, 13, 1, 1, true)
	TestCase("2017-02-02", "2018-03-01", 2, 13, 2, 1, true)
	TestCase("2017-02-02", "2018-03-01", 3, 13, 2, 1, true)
	TestCase("2017-02-02", "2018-03-01", 4, 13, 2, 1, true)
	TestCase("2017-02-02", "2018-03-01", 5, 13, 2, 1, true)
	TestCase("2017-02-02", "2018-03-01", 6, 13, 2, 1, true)
	TestCase("2017-02-02", "2018-03-01", 7, 13, 2, 1, true)
	TestCase("2017-02-02", "2018-03-01", 8, 13, 2, 1, true)
	TestCase("2017-02-02", "2018-03-01", 9, 13, 2, 1, true)
	TestCase("2017-02-02", "2018-03-01", 10, 13, 2, 1, true)
	TestCase("2017-02-02", "2018-03-01", 11, 13, 2, 1, true)
	TestCase("2017-02-02", "2018-03-01", 12, 13, 2, 1, true)
	TestCase("2017-02-02", "2018-03-01", 13, 13, 2, 1, true)
	TestCase("2017-02-02", "2018-03-01", 13, 13, 3, 1, false)

	TestCase("2017-02-02", "2018-02-15", 1, 13, 1, 1, true)
	TestCase("2017-02-02", "2018-02-15", 2, 13, 2, 1, true)
	TestCase("2017-02-02", "2018-02-15", 3, 13, 2, 1, true)
	TestCase("2017-02-02", "2018-02-15", 4, 13, 2, 1, true)
	TestCase("2017-02-02", "2018-02-15", 5, 13, 2, 1, true)
	TestCase("2017-02-02", "2018-02-15", 6, 13, 2, 1, true)
	TestCase("2017-02-02", "2018-02-15", 7, 13, 2, 1, true)
	TestCase("2017-02-02", "2018-02-15", 8, 13, 2, 1, true)
	TestCase("2017-02-02", "2018-02-15", 9, 13, 2, 1, true)
	TestCase("2017-02-02", "2018-02-15", 10, 13, 2, 1, true)
	TestCase("2017-02-02", "2018-02-15", 11, 13, 2, 1, true)
	TestCase("2017-02-02", "2018-02-15", 12, 13, 2, 1, true)
	TestCase("2017-02-02", "2018-02-15", 13, 13, 2, 1, false)
	TestCase("2017-02-02", "2018-02-15", 13, 13, 3, 1, false)

	fmt.Println("---------------------------------------------")

	TestCase("2017-04-01", "2017-05-31", 1, 1, 1, 2, true)

	TestCase("2017-01-01", "2017-02-31", 1, 1, 1, 2, true)

	TestCase("2017-01-02", "2017-09-01", 1, 4, 1, 2, true)
	TestCase("2017-01-02", "2017-09-01", 2, 4, 2, 2, true)
	TestCase("2017-01-02", "2017-09-01", 3, 4, 2, 2, true)
	TestCase("2017-01-02", "2017-09-01", 4, 4, 2, 2, true)
	TestCase("2017-01-02", "2017-09-01", 4, 4, 3, 2, false)

	TestCase("2017-11-26", "2018-01-25", 1, 1, 1, 2, true)

	TestCase("2017-11-26", "2018-01-26", 1, 2, 1, 3, false)

	TestCase("2017-03-01", "2017-10-31", 1, 4, 1, 2, true)
	TestCase("2017-03-01", "2017-10-31", 2, 4, 2, 2, true)
	TestCase("2017-03-01", "2017-10-31", 3, 4, 2, 2, true)
	TestCase("2017-03-01", "2017-10-31", 4, 4, 2, 2, true)

	TestCase("2017-02-02", "2018-03-01", 1, 7, 1, 2, true)
	TestCase("2017-02-02", "2018-03-01", 2, 7, 2, 2, true)
	TestCase("2017-02-02", "2018-03-01", 3, 7, 2, 2, true)
	TestCase("2017-02-02", "2018-03-01", 4, 7, 2, 2, true)
	TestCase("2017-02-02", "2018-03-01", 5, 7, 2, 2, true)
	TestCase("2017-02-02", "2018-03-01", 6, 7, 2, 2, true)
	TestCase("2017-02-02", "2018-03-01", 7, 7, 2, 2, false)
	TestCase("2017-02-02", "2018-03-01", 7, 7, 3, 2, false)

	TestCase("2017-02-02", "2018-02-15", 1, 5, 1, 3, true)
	TestCase("2017-02-02", "2018-02-15", 2, 5, 2, 3, true)
	TestCase("2017-02-02", "2018-02-15", 3, 5, 2, 3, true)
	TestCase("2017-02-02", "2018-02-15", 4, 5, 2, 3, true)
	TestCase("2017-02-02", "2018-02-15", 5, 5, 2, 3, false)
	TestCase("2017-02-02", "2018-02-15", 5, 5, 3, 3, false)

	fmt.Println("---------------------------------------------")

	TestCase("2017-01-28", "2018-02-15", 1, 13, 1, 1, true)
	TestCase("2017-01-28", "2018-02-15", 2, 13, 2, 1, true)
	TestCase("2017-01-28", "2018-02-15", 3, 13, 2, 1, true)
	TestCase("2017-01-28", "2018-02-15", 4, 13, 2, 1, true)
	TestCase("2017-01-28", "2018-02-15", 5, 13, 2, 1, true)
	TestCase("2017-01-28", "2018-02-15", 6, 13, 2, 1, true)
	TestCase("2017-01-28", "2018-02-15", 7, 13, 2, 1, true)
	TestCase("2017-01-28", "2018-02-15", 8, 13, 2, 1, true)
	TestCase("2017-01-28", "2018-02-15", 9, 13, 2, 1, true)
	TestCase("2017-01-28", "2018-02-15", 10, 13, 2, 1, true)
	TestCase("2017-01-28", "2018-02-15", 11, 13, 2, 1, true)
	TestCase("2017-01-28", "2018-02-15", 12, 13, 2, 1, true)
	TestCase("2017-01-28", "2018-02-15", 13, 13, 2, 1, false)
	TestCase("2017-01-28", "2018-02-15", 13, 13, 3, 1, false)

	TestCase("2017-01-28", "2018-02-15", 1, 7, 1, 2, true)
	TestCase("2017-01-28", "2018-02-15", 2, 7, 2, 2, true)
	TestCase("2017-01-28", "2018-02-15", 3, 7, 2, 2, true)
	TestCase("2017-01-28", "2018-02-15", 4, 7, 2, 2, true)
	TestCase("2017-01-28", "2018-02-15", 5, 7, 2, 2, true)
	TestCase("2017-01-28", "2018-02-15", 6, 7, 2, 2, true)
	TestCase("2017-01-28", "2018-02-15", 7, 7, 2, 2, false)
	TestCase("2017-01-28", "2018-02-15", 7, 7, 3, 2, false)

	TestCase("2017-01-28", "2018-02-15", 1, 5, 1, 3, true)
	TestCase("2017-01-28", "2018-02-15", 2, 5, 2, 3, true)
	TestCase("2017-01-28", "2018-02-15", 3, 5, 2, 3, true)
	TestCase("2017-01-28", "2018-02-15", 4, 5, 2, 3, true)
	TestCase("2017-01-28", "2018-02-15", 5, 5, 2, 3, false)
	TestCase("2017-01-28", "2018-02-15", 5, 5, 3, 3, false)

	TestCase("2017-01-28", "2018-02-15", 1, 4, 1, 4, true)
	TestCase("2017-01-28", "2018-02-15", 2, 4, 2, 4, true)
	TestCase("2017-01-28", "2018-02-15", 3, 4, 2, 4, true)
	TestCase("2017-01-28", "2018-02-15", 4, 4, 2, 4, false)
	TestCase("2017-01-28", "2018-02-15", 4, 4, 3, 4, false)

	fmt.Println("---------------------------------------------")

	TestCase("2017-01-29", "2018-02-15", 1, 13, 1, 1, true)
	TestCase("2017-01-29", "2018-02-15", 2, 13, 2, 1, true)
	TestCase("2017-01-29", "2018-02-15", 3, 13, 2, 1, true)
	TestCase("2017-01-29", "2018-02-15", 4, 13, 2, 1, true)
	TestCase("2017-01-29", "2018-02-15", 5, 13, 2, 1, true)
	TestCase("2017-01-29", "2018-02-15", 6, 13, 2, 1, true)
	TestCase("2017-01-29", "2018-02-15", 7, 13, 2, 1, true)
	TestCase("2017-01-29", "2018-02-15", 8, 13, 2, 1, true)
	TestCase("2017-01-29", "2018-02-15", 9, 13, 2, 1, true)
	TestCase("2017-01-29", "2018-02-15", 10, 13, 2, 1, true)
	TestCase("2017-01-29", "2018-02-15", 11, 13, 2, 1, true)
	TestCase("2017-01-29", "2018-02-15", 12, 13, 2, 1, true)
	TestCase("2017-01-29", "2018-02-15", 13, 13, 2, 1, false)
	TestCase("2017-01-29", "2018-02-15", 13, 13, 3, 1, false)

	TestCase("2017-01-29", "2018-02-15", 1, 7, 1, 2, true)
	TestCase("2017-01-29", "2018-02-15", 2, 7, 2, 2, true)
	TestCase("2017-01-29", "2018-02-15", 3, 7, 2, 2, true)
	TestCase("2017-01-29", "2018-02-15", 4, 7, 2, 2, true)
	TestCase("2017-01-29", "2018-02-15", 5, 7, 2, 2, true)
	TestCase("2017-01-29", "2018-02-15", 6, 7, 2, 2, true)
	TestCase("2017-01-29", "2018-02-15", 7, 7, 2, 2, false)
	TestCase("2017-01-29", "2018-02-15", 7, 7, 3, 2, false)

	TestCase("2017-01-29", "2018-02-15", 1, 5, 1, 3, true)
	TestCase("2017-01-29", "2018-02-15", 2, 5, 2, 3, true)
	TestCase("2017-01-29", "2018-02-15", 3, 5, 2, 3, true)
	TestCase("2017-01-29", "2018-02-15", 4, 5, 2, 3, true)
	TestCase("2017-01-29", "2018-02-15", 5, 5, 2, 3, false)
	TestCase("2017-01-29", "2018-02-15", 5, 5, 3, 3, false)

	TestCase("2017-01-29", "2018-02-15", 1, 4, 1, 4, true)
	TestCase("2017-01-29", "2018-02-15", 2, 4, 2, 4, true)
	TestCase("2017-01-29", "2018-02-15", 3, 4, 2, 4, true)
	TestCase("2017-01-29", "2018-02-15", 4, 4, 2, 4, false)
	TestCase("2017-01-29", "2018-02-15", 4, 4, 3, 4, false)

	fmt.Println("---------------------------------------------")

	TestCase("2017-01-30", "2018-02-15", 1, 13, 1, 1, true)
	TestCase("2017-01-30", "2018-02-15", 2, 13, 2, 1, true)
	TestCase("2017-01-30", "2018-02-15", 3, 13, 2, 1, true)
	TestCase("2017-01-30", "2018-02-15", 4, 13, 2, 1, true)
	TestCase("2017-01-30", "2018-02-15", 5, 13, 2, 1, true)
	TestCase("2017-01-30", "2018-02-15", 6, 13, 2, 1, true)
	TestCase("2017-01-30", "2018-02-15", 7, 13, 2, 1, true)
	TestCase("2017-01-30", "2018-02-15", 8, 13, 2, 1, true)
	TestCase("2017-01-30", "2018-02-15", 9, 13, 2, 1, true)
	TestCase("2017-01-30", "2018-02-15", 10, 13, 2, 1, true)
	TestCase("2017-01-30", "2018-02-15", 11, 13, 2, 1, true)
	TestCase("2017-01-30", "2018-02-15", 12, 13, 2, 1, true)
	TestCase("2017-01-30", "2018-02-15", 13, 13, 2, 1, false)
	TestCase("2017-01-30", "2018-02-15", 13, 13, 3, 1, false)

	TestCase("2017-01-30", "2018-02-15", 1, 7, 1, 2, true)
	TestCase("2017-01-30", "2018-02-15", 2, 7, 2, 2, true)
	TestCase("2017-01-30", "2018-02-15", 3, 7, 2, 2, true)
	TestCase("2017-01-30", "2018-02-15", 4, 7, 2, 2, true)
	TestCase("2017-01-30", "2018-02-15", 5, 7, 2, 2, true)
	TestCase("2017-01-30", "2018-02-15", 6, 7, 2, 2, true)
	TestCase("2017-01-30", "2018-02-15", 7, 7, 2, 2, false)
	TestCase("2017-01-30", "2018-02-15", 7, 7, 3, 2, false)

	TestCase("2017-01-30", "2018-02-15", 1, 5, 1, 3, true)
	TestCase("2017-01-30", "2018-02-15", 2, 5, 2, 3, true)
	TestCase("2017-01-30", "2018-02-15", 3, 5, 2, 3, true)
	TestCase("2017-01-30", "2018-02-15", 4, 5, 2, 3, true)
	TestCase("2017-01-30", "2018-02-15", 5, 5, 2, 3, false)
	TestCase("2017-01-30", "2018-02-15", 5, 5, 3, 3, false)

	TestCase("2017-01-30", "2018-02-15", 1, 4, 1, 4, true)
	TestCase("2017-01-30", "2018-02-15", 2, 4, 2, 4, true)
	TestCase("2017-01-30", "2018-02-15", 3, 4, 2, 4, true)
	TestCase("2017-01-30", "2018-02-15", 4, 4, 2, 4, false)
	TestCase("2017-01-30", "2018-02-15", 4, 4, 3, 4, false)

	fmt.Println("---------------------------------------------")

	TestCase("2017-01-31", "2018-02-15", 1, 13, 1, 1, true)
	TestCase("2017-01-31", "2018-02-15", 2, 13, 2, 1, true)
	TestCase("2017-01-31", "2018-02-15", 3, 13, 2, 1, true)
	TestCase("2017-01-31", "2018-02-15", 4, 13, 2, 1, true)
	TestCase("2017-01-31", "2018-02-15", 5, 13, 2, 1, true)
	TestCase("2017-01-31", "2018-02-15", 6, 13, 2, 1, true)
	TestCase("2017-01-31", "2018-02-15", 7, 13, 2, 1, true)
	TestCase("2017-01-31", "2018-02-15", 8, 13, 2, 1, true)
	TestCase("2017-01-31", "2018-02-15", 9, 13, 2, 1, true)
	TestCase("2017-01-31", "2018-02-15", 10, 13, 2, 1, true)
	TestCase("2017-01-31", "2018-02-15", 11, 13, 2, 1, true)
	TestCase("2017-01-31", "2018-02-15", 12, 13, 2, 1, true)
	TestCase("2017-01-31", "2018-02-15", 13, 13, 2, 1, false)
	TestCase("2017-01-31", "2018-02-15", 13, 13, 3, 1, false)

	TestCase("2017-01-31", "2018-02-15", 1, 7, 1, 2, true)
	TestCase("2017-01-31", "2018-02-15", 2, 7, 2, 2, true)
	TestCase("2017-01-31", "2018-02-15", 3, 7, 2, 2, true)
	TestCase("2017-01-31", "2018-02-15", 4, 7, 2, 2, true)
	TestCase("2017-01-31", "2018-02-15", 5, 7, 2, 2, true)
	TestCase("2017-01-31", "2018-02-15", 6, 7, 2, 2, true)
	TestCase("2017-01-31", "2018-02-15", 7, 7, 2, 2, false)
	TestCase("2017-01-31", "2018-02-15", 7, 7, 3, 2, false)

	TestCase("2017-01-31", "2018-02-15", 1, 5, 1, 3, true)
	TestCase("2017-01-31", "2018-02-15", 2, 5, 2, 3, true)
	TestCase("2017-01-31", "2018-02-15", 3, 5, 2, 3, true)
	TestCase("2017-01-31", "2018-02-15", 4, 5, 2, 3, true)
	TestCase("2017-01-31", "2018-02-15", 5, 5, 2, 3, false)
	TestCase("2017-01-31", "2018-02-15", 5, 5, 3, 3, false)

	TestCase("2017-01-31", "2018-02-15", 1, 4, 1, 4, true)
	TestCase("2017-01-31", "2018-02-15", 2, 4, 2, 4, true)
	TestCase("2017-01-31", "2018-02-15", 3, 4, 2, 4, true)
	TestCase("2017-01-31", "2018-02-15", 4, 4, 2, 4, false)
	TestCase("2017-01-31", "2018-02-15", 4, 4, 3, 4, false)

	fmt.Println("---------------------------------------------")

	TestCase("2017-02-01", "2018-02-15", 1, 13, 1, 1, true)
	TestCase("2017-02-01", "2018-02-15", 2, 13, 2, 1, true)
	TestCase("2017-02-01", "2018-02-15", 3, 13, 2, 1, true)
	TestCase("2017-02-01", "2018-02-15", 4, 13, 2, 1, true)
	TestCase("2017-02-01", "2018-02-15", 5, 13, 2, 1, true)
	TestCase("2017-02-01", "2018-02-15", 6, 13, 2, 1, true)
	TestCase("2017-02-01", "2018-02-15", 7, 13, 2, 1, true)
	TestCase("2017-02-01", "2018-02-15", 8, 13, 2, 1, true)
	TestCase("2017-02-01", "2018-02-15", 9, 13, 2, 1, true)
	TestCase("2017-02-01", "2018-02-15", 10, 13, 2, 1, true)
	TestCase("2017-02-01", "2018-02-15", 11, 13, 2, 1, true)
	TestCase("2017-02-01", "2018-02-15", 12, 13, 2, 1, true)
	TestCase("2017-02-01", "2018-02-15", 13, 13, 2, 1, false)
	TestCase("2017-02-01", "2018-02-15", 13, 13, 3, 1, false)

	TestCase("2017-02-01", "2018-02-15", 1, 7, 1, 2, true)
	TestCase("2017-02-01", "2018-02-15", 2, 7, 2, 2, true)
	TestCase("2017-02-01", "2018-02-15", 3, 7, 2, 2, true)
	TestCase("2017-02-01", "2018-02-15", 4, 7, 2, 2, true)
	TestCase("2017-02-01", "2018-02-15", 5, 7, 2, 2, true)
	TestCase("2017-02-01", "2018-02-15", 6, 7, 2, 2, true)
	TestCase("2017-02-01", "2018-02-15", 7, 7, 2, 2, false)
	TestCase("2017-02-01", "2018-02-15", 7, 7, 3, 2, false)

	TestCase("2017-02-01", "2018-02-15", 1, 5, 1, 3, true)
	TestCase("2017-02-01", "2018-02-15", 2, 5, 2, 3, true)
	TestCase("2017-02-01", "2018-02-15", 3, 5, 2, 3, true)
	TestCase("2017-02-01", "2018-02-15", 4, 5, 2, 3, true)
	TestCase("2017-02-01", "2018-02-15", 5, 5, 2, 3, false)
	TestCase("2017-02-01", "2018-02-15", 5, 5, 3, 3, false)

	TestCase("2017-02-01", "2018-02-15", 1, 4, 1, 4, true)
	TestCase("2017-02-01", "2018-02-15", 2, 4, 2, 4, true)
	TestCase("2017-02-01", "2018-02-15", 3, 4, 2, 4, true)
	TestCase("2017-02-01", "2018-02-15", 4, 4, 2, 4, false)
	TestCase("2017-02-01", "2018-02-15", 4, 4, 3, 4, false)

	fmt.Println("---------------------------------------------")
	fmt.Println("---------------------------------------------")

	TestCase("2020-04-01", "2020-05-31", 1, 2, 1, 1, true)
	TestCase("2020-04-01", "2020-05-31", 2, 2, 2, 1, true)

	TestCase("2020-01-01", "2020-02-31", 1, 2, 1, 1, true)
	TestCase("2020-01-01", "2020-02-31", 2, 2, 2, 1, true)

	TestCase("2020-01-02", "2020-09-01", 1, 8, 1, 1, true)
	TestCase("2020-01-02", "2020-09-01", 2, 8, 2, 1, true)
	TestCase("2020-01-02", "2020-09-01", 3, 8, 2, 1, true)
	TestCase("2020-01-02", "2020-09-01", 4, 8, 2, 1, true)
	TestCase("2020-01-02", "2020-09-01", 5, 8, 2, 1, true)
	TestCase("2020-01-02", "2020-09-01", 6, 8, 2, 1, true)
	TestCase("2020-01-02", "2020-09-01", 7, 8, 2, 1, true)
	TestCase("2020-01-02", "2020-09-01", 8, 8, 2, 1, true)
	TestCase("2020-01-02", "2020-09-01", 8, 8, 3, 1, false)

	TestCase("2020-11-26", "2021-01-25", 1, 2, 1, 1, true)
	TestCase("2020-11-26", "2021-01-25", 2, 2, 2, 1, true)

	TestCase("2020-11-26", "2021-01-26", 1, 3, 1, 1, true)
	TestCase("2020-11-26", "2021-01-26", 2, 3, 2, 1, true)
	TestCase("2020-11-26", "2021-01-26", 3, 3, 2, 1, false)

	TestCase("2020-03-01", "2020-10-31", 1, 8, 1, 1, true)
	TestCase("2020-03-01", "2020-10-31", 2, 8, 2, 1, true)
	TestCase("2020-03-01", "2020-10-31", 3, 8, 2, 1, true)
	TestCase("2020-03-01", "2020-10-31", 4, 8, 2, 1, true)
	TestCase("2020-03-01", "2020-10-31", 5, 8, 2, 1, true)
	TestCase("2020-03-01", "2020-10-31", 6, 8, 2, 1, true)
	TestCase("2020-03-01", "2020-10-31", 7, 8, 2, 1, true)
	TestCase("2020-03-01", "2020-10-31", 8, 8, 2, 1, true)

	TestCase("2020-02-02", "2021-03-01", 1, 13, 1, 1, true)
	TestCase("2020-02-02", "2021-03-01", 2, 13, 2, 1, true)
	TestCase("2020-02-02", "2021-03-01", 3, 13, 2, 1, true)
	TestCase("2020-02-02", "2021-03-01", 4, 13, 2, 1, true)
	TestCase("2020-02-02", "2021-03-01", 5, 13, 2, 1, true)
	TestCase("2020-02-02", "2021-03-01", 6, 13, 2, 1, true)
	TestCase("2020-02-02", "2021-03-01", 7, 13, 2, 1, true)
	TestCase("2020-02-02", "2021-03-01", 8, 13, 2, 1, true)
	TestCase("2020-02-02", "2021-03-01", 9, 13, 2, 1, true)
	TestCase("2020-02-02", "2021-03-01", 10, 13, 2, 1, true)
	TestCase("2020-02-02", "2021-03-01", 11, 13, 2, 1, true)
	TestCase("2020-02-02", "2021-03-01", 12, 13, 2, 1, true)
	TestCase("2020-02-02", "2021-03-01", 13, 13, 2, 1, true)
	TestCase("2020-02-02", "2021-03-01", 13, 13, 3, 1, false)

	TestCase("2020-02-02", "2021-02-15", 1, 13, 1, 1, true)
	TestCase("2020-02-02", "2021-02-15", 2, 13, 2, 1, true)
	TestCase("2020-02-02", "2021-02-15", 3, 13, 2, 1, true)
	TestCase("2020-02-02", "2021-02-15", 4, 13, 2, 1, true)
	TestCase("2020-02-02", "2021-02-15", 5, 13, 2, 1, true)
	TestCase("2020-02-02", "2021-02-15", 6, 13, 2, 1, true)
	TestCase("2020-02-02", "2021-02-15", 7, 13, 2, 1, true)
	TestCase("2020-02-02", "2021-02-15", 8, 13, 2, 1, true)
	TestCase("2020-02-02", "2021-02-15", 9, 13, 2, 1, true)
	TestCase("2020-02-02", "2021-02-15", 10, 13, 2, 1, true)
	TestCase("2020-02-02", "2021-02-15", 11, 13, 2, 1, true)
	TestCase("2020-02-02", "2021-02-15", 12, 13, 2, 1, true)
	TestCase("2020-02-02", "2021-02-15", 13, 13, 2, 1, false)
	TestCase("2020-02-02", "2021-02-15", 13, 13, 3, 1, false)

	fmt.Println("---------------------------------------------")

	TestCase("2020-04-01", "2020-05-31", 1, 1, 1, 2, true)

	TestCase("2020-01-01", "2020-02-31", 1, 1, 1, 2, true)

	TestCase("2020-01-02", "2020-09-01", 1, 4, 1, 2, true)
	TestCase("2020-01-02", "2020-09-01", 2, 4, 2, 2, true)
	TestCase("2020-01-02", "2020-09-01", 3, 4, 2, 2, true)
	TestCase("2020-01-02", "2020-09-01", 4, 4, 2, 2, true)
	TestCase("2020-01-02", "2020-09-01", 4, 4, 3, 2, false)

	TestCase("2020-11-26", "2021-01-25", 1, 1, 1, 2, true)

	TestCase("2020-11-26", "2021-01-26", 1, 2, 1, 3, false)

	TestCase("2020-03-01", "2020-10-31", 1, 4, 1, 2, true)
	TestCase("2020-03-01", "2020-10-31", 2, 4, 2, 2, true)
	TestCase("2020-03-01", "2020-10-31", 3, 4, 2, 2, true)
	TestCase("2020-03-01", "2020-10-31", 4, 4, 2, 2, true)

	TestCase("2020-02-02", "2021-03-01", 1, 7, 1, 2, true)
	TestCase("2020-02-02", "2021-03-01", 2, 7, 2, 2, true)
	TestCase("2020-02-02", "2021-03-01", 3, 7, 2, 2, true)
	TestCase("2020-02-02", "2021-03-01", 4, 7, 2, 2, true)
	TestCase("2020-02-02", "2021-03-01", 5, 7, 2, 2, true)
	TestCase("2020-02-02", "2021-03-01", 6, 7, 2, 2, true)
	TestCase("2020-02-02", "2021-03-01", 7, 7, 2, 2, false)
	TestCase("2020-02-02", "2021-03-01", 7, 7, 3, 2, false)

	TestCase("2020-02-02", "2021-02-15", 1, 5, 1, 3, true)
	TestCase("2020-02-02", "2021-02-15", 2, 5, 2, 3, true)
	TestCase("2020-02-02", "2021-02-15", 3, 5, 2, 3, true)
	TestCase("2020-02-02", "2021-02-15", 4, 5, 2, 3, true)
	TestCase("2020-02-02", "2021-02-15", 5, 5, 2, 3, false)
	TestCase("2020-02-02", "2021-02-15", 5, 5, 3, 3, false)

	fmt.Println("---------------------------------------------")

	TestCase("2020-01-28", "2021-02-15", 1, 13, 1, 1, true)
	TestCase("2020-01-28", "2021-02-15", 2, 13, 2, 1, true)
	TestCase("2020-01-28", "2021-02-15", 3, 13, 2, 1, true)
	TestCase("2020-01-28", "2021-02-15", 4, 13, 2, 1, true)
	TestCase("2020-01-28", "2021-02-15", 5, 13, 2, 1, true)
	TestCase("2020-01-28", "2021-02-15", 6, 13, 2, 1, true)
	TestCase("2020-01-28", "2021-02-15", 7, 13, 2, 1, true)
	TestCase("2020-01-28", "2021-02-15", 8, 13, 2, 1, true)
	TestCase("2020-01-28", "2021-02-15", 9, 13, 2, 1, true)
	TestCase("2020-01-28", "2021-02-15", 10, 13, 2, 1, true)
	TestCase("2020-01-28", "2021-02-15", 11, 13, 2, 1, true)
	TestCase("2020-01-28", "2021-02-15", 12, 13, 2, 1, true)
	TestCase("2020-01-28", "2021-02-15", 13, 13, 2, 1, false)
	TestCase("2020-01-28", "2021-02-15", 13, 13, 3, 1, false)

	TestCase("2020-01-28", "2021-02-15", 1, 7, 1, 2, true)
	TestCase("2020-01-28", "2021-02-15", 2, 7, 2, 2, true)
	TestCase("2020-01-28", "2021-02-15", 3, 7, 2, 2, true)
	TestCase("2020-01-28", "2021-02-15", 4, 7, 2, 2, true)
	TestCase("2020-01-28", "2021-02-15", 5, 7, 2, 2, true)
	TestCase("2020-01-28", "2021-02-15", 6, 7, 2, 2, true)
	TestCase("2020-01-28", "2021-02-15", 7, 7, 2, 2, false)
	TestCase("2020-01-28", "2021-02-15", 7, 7, 3, 2, false)

	TestCase("2020-01-28", "2021-02-15", 1, 5, 1, 3, true)
	TestCase("2020-01-28", "2021-02-15", 2, 5, 2, 3, true)
	TestCase("2020-01-28", "2021-02-15", 3, 5, 2, 3, true)
	TestCase("2020-01-28", "2021-02-15", 4, 5, 2, 3, true)
	TestCase("2020-01-28", "2021-02-15", 5, 5, 2, 3, false)
	TestCase("2020-01-28", "2021-02-15", 5, 5, 3, 3, false)

	TestCase("2020-01-28", "2021-02-15", 1, 4, 1, 4, true)
	TestCase("2020-01-28", "2021-02-15", 2, 4, 2, 4, true)
	TestCase("2020-01-28", "2021-02-15", 3, 4, 2, 4, true)
	TestCase("2020-01-28", "2021-02-15", 4, 4, 2, 4, false)
	TestCase("2020-01-28", "2021-02-15", 4, 4, 3, 4, false)

	fmt.Println("---------------------------------------------")

	TestCase("2020-01-29", "2021-02-15", 1, 13, 1, 1, true)
	TestCase("2020-01-29", "2021-02-15", 2, 13, 2, 1, true)
	TestCase("2020-01-29", "2021-02-15", 3, 13, 2, 1, true)
	TestCase("2020-01-29", "2021-02-15", 4, 13, 2, 1, true)
	TestCase("2020-01-29", "2021-02-15", 5, 13, 2, 1, true)
	TestCase("2020-01-29", "2021-02-15", 6, 13, 2, 1, true)
	TestCase("2020-01-29", "2021-02-15", 7, 13, 2, 1, true)
	TestCase("2020-01-29", "2021-02-15", 8, 13, 2, 1, true)
	TestCase("2020-01-29", "2021-02-15", 9, 13, 2, 1, true)
	TestCase("2020-01-29", "2021-02-15", 10, 13, 2, 1, true)
	TestCase("2020-01-29", "2021-02-15", 11, 13, 2, 1, true)
	TestCase("2020-01-29", "2021-02-15", 12, 13, 2, 1, true)
	TestCase("2020-01-29", "2021-02-15", 13, 13, 2, 1, false)
	TestCase("2020-01-29", "2021-02-15", 13, 13, 3, 1, false)

	TestCase("2020-01-29", "2021-02-15", 1, 7, 1, 2, true)
	TestCase("2020-01-29", "2021-02-15", 2, 7, 2, 2, true)
	TestCase("2020-01-29", "2021-02-15", 3, 7, 2, 2, true)
	TestCase("2020-01-29", "2021-02-15", 4, 7, 2, 2, true)
	TestCase("2020-01-29", "2021-02-15", 5, 7, 2, 2, true)
	TestCase("2020-01-29", "2021-02-15", 6, 7, 2, 2, true)
	TestCase("2020-01-29", "2021-02-15", 7, 7, 2, 2, false)
	TestCase("2020-01-29", "2021-02-15", 7, 7, 3, 2, false)

	TestCase("2020-01-29", "2021-02-15", 1, 5, 1, 3, true)
	TestCase("2020-01-29", "2021-02-15", 2, 5, 2, 3, true)
	TestCase("2020-01-29", "2021-02-15", 3, 5, 2, 3, true)
	TestCase("2020-01-29", "2021-02-15", 4, 5, 2, 3, true)
	TestCase("2020-01-29", "2021-02-15", 5, 5, 2, 3, false)
	TestCase("2020-01-29", "2021-02-15", 5, 5, 3, 3, false)

	TestCase("2020-01-29", "2021-02-15", 1, 4, 1, 4, true)
	TestCase("2020-01-29", "2021-02-15", 2, 4, 2, 4, true)
	TestCase("2020-01-29", "2021-02-15", 3, 4, 2, 4, true)
	TestCase("2020-01-29", "2021-02-15", 4, 4, 2, 4, false)
	TestCase("2020-01-29", "2021-02-15", 4, 4, 3, 4, false)

	fmt.Println("---------------------------------------------")

	TestCase("2020-01-30", "2021-02-15", 1, 13, 1, 1, true)
	TestCase("2020-01-30", "2021-02-15", 2, 13, 2, 1, true)
	TestCase("2020-01-30", "2021-02-15", 3, 13, 2, 1, true)
	TestCase("2020-01-30", "2021-02-15", 4, 13, 2, 1, true)
	TestCase("2020-01-30", "2021-02-15", 5, 13, 2, 1, true)
	TestCase("2020-01-30", "2021-02-15", 6, 13, 2, 1, true)
	TestCase("2020-01-30", "2021-02-15", 7, 13, 2, 1, true)
	TestCase("2020-01-30", "2021-02-15", 8, 13, 2, 1, true)
	TestCase("2020-01-30", "2021-02-15", 9, 13, 2, 1, true)
	TestCase("2020-01-30", "2021-02-15", 10, 13, 2, 1, true)
	TestCase("2020-01-30", "2021-02-15", 11, 13, 2, 1, true)
	TestCase("2020-01-30", "2021-02-15", 12, 13, 2, 1, true)
	TestCase("2020-01-30", "2021-02-15", 13, 13, 2, 1, false)
	TestCase("2020-01-30", "2021-02-15", 13, 13, 3, 1, false)

	TestCase("2020-01-30", "2021-02-15", 1, 7, 1, 2, true)
	TestCase("2020-01-30", "2021-02-15", 2, 7, 2, 2, true)
	TestCase("2020-01-30", "2021-02-15", 3, 7, 2, 2, true)
	TestCase("2020-01-30", "2021-02-15", 4, 7, 2, 2, true)
	TestCase("2020-01-30", "2021-02-15", 5, 7, 2, 2, true)
	TestCase("2020-01-30", "2021-02-15", 6, 7, 2, 2, true)
	TestCase("2020-01-30", "2021-02-15", 7, 7, 2, 2, false)
	TestCase("2020-01-30", "2021-02-15", 7, 7, 3, 2, false)

	TestCase("2020-01-30", "2021-02-15", 1, 5, 1, 3, true)
	TestCase("2020-01-30", "2021-02-15", 2, 5, 2, 3, true)
	TestCase("2020-01-30", "2021-02-15", 3, 5, 2, 3, true)
	TestCase("2020-01-30", "2021-02-15", 4, 5, 2, 3, true)
	TestCase("2020-01-30", "2021-02-15", 5, 5, 2, 3, false)
	TestCase("2020-01-30", "2021-02-15", 5, 5, 3, 3, false)

	TestCase("2020-01-30", "2021-02-15", 1, 4, 1, 4, true)
	TestCase("2020-01-30", "2021-02-15", 2, 4, 2, 4, true)
	TestCase("2020-01-30", "2021-02-15", 3, 4, 2, 4, true)
	TestCase("2020-01-30", "2021-02-15", 4, 4, 2, 4, false)
	TestCase("2020-01-30", "2021-02-15", 4, 4, 3, 4, false)

	fmt.Println("---------------------------------------------")

	TestCase("2020-01-31", "2021-02-15", 1, 13, 1, 1, true)
	TestCase("2020-01-31", "2021-02-15", 2, 13, 2, 1, true)
	TestCase("2020-01-31", "2021-02-15", 3, 13, 2, 1, true)
	TestCase("2020-01-31", "2021-02-15", 4, 13, 2, 1, true)
	TestCase("2020-01-31", "2021-02-15", 5, 13, 2, 1, true)
	TestCase("2020-01-31", "2021-02-15", 6, 13, 2, 1, true)
	TestCase("2020-01-31", "2021-02-15", 7, 13, 2, 1, true)
	TestCase("2020-01-31", "2021-02-15", 8, 13, 2, 1, true)
	TestCase("2020-01-31", "2021-02-15", 9, 13, 2, 1, true)
	TestCase("2020-01-31", "2021-02-15", 10, 13, 2, 1, true)
	TestCase("2020-01-31", "2021-02-15", 11, 13, 2, 1, true)
	TestCase("2020-01-31", "2021-02-15", 12, 13, 2, 1, true)
	TestCase("2020-01-31", "2021-02-15", 13, 13, 2, 1, false)
	TestCase("2020-01-31", "2021-02-15", 13, 13, 3, 1, false)

	TestCase("2020-01-31", "2021-02-15", 1, 7, 1, 2, true)
	TestCase("2020-01-31", "2021-02-15", 2, 7, 2, 2, true)
	TestCase("2020-01-31", "2021-02-15", 3, 7, 2, 2, true)
	TestCase("2020-01-31", "2021-02-15", 4, 7, 2, 2, true)
	TestCase("2020-01-31", "2021-02-15", 5, 7, 2, 2, true)
	TestCase("2020-01-31", "2021-02-15", 6, 7, 2, 2, true)
	TestCase("2020-01-31", "2021-02-15", 7, 7, 2, 2, false)
	TestCase("2020-01-31", "2021-02-15", 7, 7, 3, 2, false)

	TestCase("2020-01-31", "2021-02-15", 1, 5, 1, 3, true)
	TestCase("2020-01-31", "2021-02-15", 2, 5, 2, 3, true)
	TestCase("2020-01-31", "2021-02-15", 3, 5, 2, 3, true)
	TestCase("2020-01-31", "2021-02-15", 4, 5, 2, 3, true)
	TestCase("2020-01-31", "2021-02-15", 5, 5, 2, 3, false)
	TestCase("2020-01-31", "2021-02-15", 5, 5, 3, 3, false)

	TestCase("2020-01-31", "2021-02-15", 1, 4, 1, 4, true)
	TestCase("2020-01-31", "2021-02-15", 2, 4, 2, 4, true)
	TestCase("2020-01-31", "2021-02-15", 3, 4, 2, 4, true)
	TestCase("2020-01-31", "2021-02-15", 4, 4, 2, 4, false)
	TestCase("2020-01-31", "2021-02-15", 4, 4, 3, 4, false)

	fmt.Println("---------------------------------------------")

	TestCase("2020-02-01", "2021-02-15", 1, 13, 1, 1, true)
	TestCase("2020-02-01", "2021-02-15", 2, 13, 2, 1, true)
	TestCase("2020-02-01", "2021-02-15", 3, 13, 2, 1, true)
	TestCase("2020-02-01", "2021-02-15", 4, 13, 2, 1, true)
	TestCase("2020-02-01", "2021-02-15", 5, 13, 2, 1, true)
	TestCase("2020-02-01", "2021-02-15", 6, 13, 2, 1, true)
	TestCase("2020-02-01", "2021-02-15", 7, 13, 2, 1, true)
	TestCase("2020-02-01", "2021-02-15", 8, 13, 2, 1, true)
	TestCase("2020-02-01", "2021-02-15", 9, 13, 2, 1, true)
	TestCase("2020-02-01", "2021-02-15", 10, 13, 2, 1, true)
	TestCase("2020-02-01", "2021-02-15", 11, 13, 2, 1, true)
	TestCase("2020-02-01", "2021-02-15", 12, 13, 2, 1, true)
	TestCase("2020-02-01", "2021-02-15", 13, 13, 2, 1, false)
	TestCase("2020-02-01", "2021-02-15", 13, 13, 3, 1, false)

	TestCase("2020-02-01", "2021-02-15", 1, 7, 1, 2, true)
	TestCase("2020-02-01", "2021-02-15", 2, 7, 2, 2, true)
	TestCase("2020-02-01", "2021-02-15", 3, 7, 2, 2, true)
	TestCase("2020-02-01", "2021-02-15", 4, 7, 2, 2, true)
	TestCase("2020-02-01", "2021-02-15", 5, 7, 2, 2, true)
	TestCase("2020-02-01", "2021-02-15", 6, 7, 2, 2, true)
	TestCase("2020-02-01", "2021-02-15", 7, 7, 2, 2, false)
	TestCase("2020-02-01", "2021-02-15", 7, 7, 3, 2, false)

	TestCase("2020-02-01", "2021-02-15", 1, 5, 1, 3, true)
	TestCase("2020-02-01", "2021-02-15", 2, 5, 2, 3, true)
	TestCase("2020-02-01", "2021-02-15", 3, 5, 2, 3, true)
	TestCase("2020-02-01", "2021-02-15", 4, 5, 2, 3, true)
	TestCase("2020-02-01", "2021-02-15", 5, 5, 2, 3, false)
	TestCase("2020-02-01", "2021-02-15", 5, 5, 3, 3, false)

	TestCase("2020-02-01", "2021-02-15", 1, 4, 1, 4, true)
	TestCase("2020-02-01", "2021-02-15", 2, 4, 2, 4, true)
	TestCase("2020-02-01", "2021-02-15", 3, 4, 2, 4, true)
	TestCase("2020-02-01", "2021-02-15", 4, 4, 2, 4, false)
	TestCase("2020-02-01", "2021-02-15", 4, 4, 3, 4, false)

	fmt.Println("---------------------------------------------")
	fmt.Println("---------------------------------------------")

	TestCase2("2020-04-01", "2020-06-30", 1, 2, 1, 1, true)
	TestCase2("2020-04-01", "2020-06-30", 2, 2, 2, 1, true)

	TestCase2("2017-12-28", "2018-04-30", 1, 4, 1, 1, true)
	TestCase2("2017-12-28", "2018-04-30", 2, 4, 2, 1, true)
	TestCase2("2017-12-28", "2018-04-30", 3, 4, 2, 1, true)
	TestCase2("2017-12-28", "2018-04-30", 4, 4, 2, 1, true)

	TestCase("2018-01-30", "2018-04-29", 1, 3, 1, 1, true)
	TestCase("2018-01-30", "2018-04-29", 2, 3, 2, 1, true)
	TestCase("2018-01-30", "2018-04-29", 3, 3, 2, 1, false)
}

func TestTestCalcStartEndTimeStr() {
	fmt.Println("TestTestCalcStartEndTimeStr begining...")
	fmt.Println("TestTestCalcStartEndTimeStr begining...")
	fmt.Println("TestTestCalcStartEndTimeStr begining...")
	fmt.Println("TestTestCalcStartEndTimeStr begining...")
	fmt.Println("TestTestCalcStartEndTimeStr begining...")
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 1, 13, 1, 1, 1, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 2, 13, 2, 1, 1, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 3, 13, 2, 1, 1, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 4, 13, 2, 1, 1, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 5, 13, 2, 1, 1, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 6, 13, 2, 1, 1, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 7, 13, 2, 1, 1, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 8, 13, 2, 1, 1, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 9, 13, 2, 1, 1, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 10, 13, 2, 1, 1, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 11, 13, 2, 1, 1, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 12, 13, 2, 1, 1, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 13, 13, 2, 1, 1, 0, false)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 13, 13, 3, 1, 1, 0, true)

	TestCalcStartEndTimeStr("2017-04-01", "2017-05-31", 1, 2, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-04-01", "2017-05-31", 2, 2, 2, 1, 0, 0, true)

	TestCalcStartEndTimeStr("2017-01-01", "2017-02-31", 1, 2, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-01", "2017-02-31", 2, 2, 2, 1, 0, 0, true)

	TestCalcStartEndTimeStr("2017-01-02", "2017-09-01", 1, 8, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-02", "2017-09-01", 2, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-02", "2017-09-01", 3, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-02", "2017-09-01", 4, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-02", "2017-09-01", 5, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-02", "2017-09-01", 6, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-02", "2017-09-01", 7, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-02", "2017-09-01", 8, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-02", "2017-09-01", 8, 8, 3, 1, 0, 0, false)

	TestCalcStartEndTimeStr("2017-11-26", "2018-01-25", 1, 2, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-11-26", "2018-01-25", 2, 2, 2, 1, 0, 0, true)

	TestCalcStartEndTimeStr("2017-11-26", "2018-01-26", 1, 3, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-11-26", "2018-01-26", 2, 3, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-11-26", "2018-01-26", 3, 3, 2, 1, 0, 0, false)

	TestCalcStartEndTimeStr("2017-03-01", "2017-10-31", 1, 8, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-03-01", "2017-10-31", 2, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-03-01", "2017-10-31", 3, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-03-01", "2017-10-31", 4, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-03-01", "2017-10-31", 5, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-03-01", "2017-10-31", 6, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-03-01", "2017-10-31", 7, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-03-01", "2017-10-31", 8, 8, 2, 1, 0, 0, true)

	TestCalcStartEndTimeStr("2017-02-02", "2018-03-01", 1, 13, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-03-01", 2, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-03-01", 3, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-03-01", 4, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-03-01", 5, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-03-01", 6, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-03-01", 7, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-03-01", 8, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-03-01", 9, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-03-01", 10, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-03-01", 11, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-03-01", 12, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-03-01", 13, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-03-01", 13, 13, 3, 1, 0, 0, false)

	TestCalcStartEndTimeStr("2017-02-02", "2018-02-15", 1, 13, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-02-15", 2, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-02-15", 3, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-02-15", 4, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-02-15", 5, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-02-15", 6, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-02-15", 7, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-02-15", 8, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-02-15", 9, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-02-15", 10, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-02-15", 11, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-02-15", 12, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-02-15", 13, 13, 2, 1, 0, 0, false)
	TestCalcStartEndTimeStr("2017-02-02", "2018-02-15", 13, 13, 3, 1, 0, 0, false)

	fmt.Println("---------------------------------------------")

	TestCalcStartEndTimeStr("2017-04-01", "2017-05-31", 1, 1, 1, 2, 0, 0, true)

	TestCalcStartEndTimeStr("2017-01-01", "2017-02-31", 1, 1, 1, 2, 0, 0, true)

	TestCalcStartEndTimeStr("2017-01-02", "2017-09-01", 1, 4, 1, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-02", "2017-09-01", 2, 4, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-02", "2017-09-01", 3, 4, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-02", "2017-09-01", 4, 4, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-02", "2017-09-01", 4, 4, 3, 2, 0, 0, false)

	TestCalcStartEndTimeStr("2017-11-26", "2018-01-25", 1, 1, 1, 2, 0, 0, true)

	TestCalcStartEndTimeStr("2017-11-26", "2018-01-26", 1, 1, 1, 3, 0, 0, false)

	TestCalcStartEndTimeStr("2017-03-01", "2017-10-31", 1, 4, 1, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-03-01", "2017-10-31", 2, 4, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-03-01", "2017-10-31", 3, 4, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-03-01", "2017-10-31", 4, 4, 2, 2, 0, 0, true)

	TestCalcStartEndTimeStr("2017-02-02", "2018-03-01", 1, 7, 1, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-03-01", 2, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-03-01", 3, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-03-01", 4, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-03-01", 5, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-03-01", 6, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-03-01", 7, 7, 2, 2, 0, 0, false)
	TestCalcStartEndTimeStr("2017-02-02", "2018-03-01", 7, 7, 3, 2, 0, 0, false)

	TestCalcStartEndTimeStr("2017-02-02", "2018-02-15", 1, 5, 1, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-02-15", 2, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-02-15", 3, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-02-15", 4, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-02", "2018-02-15", 5, 5, 2, 3, 0, 0, false)
	TestCalcStartEndTimeStr("2017-02-02", "2018-02-15", 5, 5, 3, 3, 0, 0, false)

	fmt.Println("---------------------------------------------")

	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 1, 13, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 2, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 3, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 4, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 5, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 6, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 7, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 8, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 9, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 10, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 11, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 12, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 13, 13, 2, 1, 0, 0, false)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 13, 13, 3, 1, 0, 0, false)

	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 1, 7, 1, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 2, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 3, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 4, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 5, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 6, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 7, 7, 2, 2, 0, 0, false)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 7, 7, 3, 2, 0, 0, false)

	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 1, 5, 1, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 2, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 3, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 4, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 5, 5, 2, 3, 0, 0, false)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 5, 5, 3, 3, 0, 0, false)

	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 1, 4, 1, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 2, 4, 2, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 3, 4, 2, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 4, 4, 2, 4, 0, 0, false)
	TestCalcStartEndTimeStr("2017-01-28", "2018-02-15", 4, 4, 3, 4, 0, 0, false)

	fmt.Println("---------------------------------------------")

	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 1, 13, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 2, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 3, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 4, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 5, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 6, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 7, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 8, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 9, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 10, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 11, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 12, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 13, 13, 2, 1, 0, 0, false)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 13, 13, 3, 1, 0, 0, false)

	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 1, 7, 1, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 2, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 3, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 4, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 5, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 6, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 7, 7, 2, 2, 0, 0, false)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 7, 7, 3, 2, 0, 0, false)

	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 1, 5, 1, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 2, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 3, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 4, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 5, 5, 2, 3, 0, 0, false)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 5, 5, 3, 3, 0, 0, false)

	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 1, 4, 1, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 2, 4, 2, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 3, 4, 2, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 4, 4, 2, 4, 0, 0, false)
	TestCalcStartEndTimeStr("2017-01-29", "2018-02-15", 4, 4, 3, 4, 0, 0, false)

	fmt.Println("---------------------------------------------")

	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 1, 13, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 2, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 3, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 4, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 5, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 6, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 7, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 8, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 9, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 10, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 11, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 12, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 13, 13, 2, 1, 0, 0, false)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 13, 13, 3, 1, 0, 0, false)

	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 1, 7, 1, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 2, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 3, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 4, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 5, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 6, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 7, 7, 2, 2, 0, 0, false)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 7, 7, 3, 2, 0, 0, false)

	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 1, 5, 1, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 2, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 3, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 4, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 5, 5, 2, 3, 0, 0, false)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 5, 5, 3, 3, 0, 0, false)

	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 1, 4, 1, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 2, 4, 2, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 3, 4, 2, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 4, 4, 2, 4, 0, 0, false)
	TestCalcStartEndTimeStr("2017-01-30", "2018-02-15", 4, 4, 3, 4, 0, 0, false)

	fmt.Println("---------------------------------------------")
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 1, 13, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 2, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 3, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 4, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 5, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 6, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 7, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 8, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 9, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 10, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 11, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 12, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 13, 13, 2, 1, 0, 0, false)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 13, 13, 3, 1, 0, 0, false)

	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 1, 7, 1, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 2, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 3, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 4, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 5, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 6, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 7, 7, 2, 2, 0, 0, false)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 7, 7, 3, 2, 0, 0, false)

	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 1, 5, 1, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 2, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 3, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 4, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 5, 5, 2, 3, 0, 0, false)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 5, 5, 3, 3, 0, 0, false)

	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 1, 4, 1, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 2, 4, 2, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 3, 4, 2, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 4, 4, 2, 4, 0, 0, false)
	TestCalcStartEndTimeStr("2017-01-31", "2018-02-15", 4, 4, 3, 4, 0, 0, false)

	fmt.Println("---------------------------------------------")

	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 1, 13, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 2, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 3, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 4, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 5, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 6, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 7, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 8, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 9, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 10, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 11, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 12, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 13, 13, 2, 1, 0, 0, false)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 13, 13, 3, 1, 0, 0, false)

	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 1, 7, 1, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 2, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 3, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 4, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 5, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 6, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 7, 7, 2, 2, 0, 0, false)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 7, 7, 3, 2, 0, 0, false)

	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 1, 5, 1, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 2, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 3, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 4, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 5, 5, 2, 3, 0, 0, false)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 5, 5, 3, 3, 0, 0, false)

	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 1, 4, 1, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 2, 4, 2, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 3, 4, 2, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 4, 4, 2, 4, 0, 0, false)
	TestCalcStartEndTimeStr("2017-02-01", "2018-02-15", 4, 4, 3, 4, 0, 0, false)

	fmt.Println("---------------------------------------------")
	fmt.Println("---------------------------------------------")

	TestCalcStartEndTimeStr("2020-04-01", "2020-05-31", 1, 2, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-04-01", "2020-05-31", 2, 2, 2, 1, 0, 0, true)

	TestCalcStartEndTimeStr("2020-01-01", "2020-02-31", 1, 2, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-01", "2020-02-31", 2, 2, 2, 1, 0, 0, true)

	TestCalcStartEndTimeStr("2020-01-02", "2020-09-01", 1, 8, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-02", "2020-09-01", 2, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-02", "2020-09-01", 3, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-02", "2020-09-01", 4, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-02", "2020-09-01", 5, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-02", "2020-09-01", 6, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-02", "2020-09-01", 7, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-02", "2020-09-01", 8, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-02", "2020-09-01", 8, 8, 3, 1, 0, 0, false)

	TestCalcStartEndTimeStr("2020-11-26", "2021-01-25", 1, 2, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-11-26", "2021-01-25", 2, 2, 2, 1, 0, 0, true)

	TestCalcStartEndTimeStr("2020-11-26", "2021-01-26", 1, 3, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-11-26", "2021-01-26", 2, 3, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-11-26", "2021-01-26", 3, 3, 2, 1, 0, 0, false)

	TestCalcStartEndTimeStr("2020-03-01", "2020-10-31", 1, 8, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-03-01", "2020-10-31", 2, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-03-01", "2020-10-31", 3, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-03-01", "2020-10-31", 4, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-03-01", "2020-10-31", 5, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-03-01", "2020-10-31", 6, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-03-01", "2020-10-31", 7, 8, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-03-01", "2020-10-31", 8, 8, 2, 1, 0, 0, true)

	TestCalcStartEndTimeStr("2020-02-02", "2021-03-01", 1, 13, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-03-01", 2, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-03-01", 3, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-03-01", 4, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-03-01", 5, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-03-01", 6, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-03-01", 7, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-03-01", 8, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-03-01", 9, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-03-01", 10, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-03-01", 11, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-03-01", 12, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-03-01", 13, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-03-01", 13, 13, 3, 1, 0, 0, false)

	TestCalcStartEndTimeStr("2020-02-02", "2021-02-15", 1, 13, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-02-15", 2, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-02-15", 3, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-02-15", 4, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-02-15", 5, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-02-15", 6, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-02-15", 7, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-02-15", 8, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-02-15", 9, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-02-15", 10, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-02-15", 11, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-02-15", 12, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-02-15", 13, 13, 2, 1, 0, 0, false)
	TestCalcStartEndTimeStr("2020-02-02", "2021-02-15", 13, 13, 3, 1, 0, 0, false)

	fmt.Println("---------------------------------------------")

	TestCalcStartEndTimeStr("2020-04-01", "2020-05-31", 1, 1, 1, 2, 0, 0, true)

	TestCalcStartEndTimeStr("2020-01-01", "2020-02-31", 1, 1, 1, 2, 0, 0, true)

	TestCalcStartEndTimeStr("2020-01-02", "2020-09-01", 1, 4, 1, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-02", "2020-09-01", 2, 4, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-02", "2020-09-01", 3, 4, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-02", "2020-09-01", 4, 4, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-02", "2020-09-01", 4, 4, 3, 2, 0, 0, false)

	TestCalcStartEndTimeStr("2020-11-26", "2021-01-25", 1, 1, 1, 2, 0, 0, true)

	TestCalcStartEndTimeStr("2020-11-26", "2021-01-26", 1, 2, 1, 3, 0, 0, false)

	TestCalcStartEndTimeStr("2020-03-01", "2020-10-31", 1, 4, 1, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-03-01", "2020-10-31", 2, 4, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-03-01", "2020-10-31", 3, 4, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-03-01", "2020-10-31", 4, 4, 2, 2, 0, 0, true)

	TestCalcStartEndTimeStr("2020-02-02", "2021-03-01", 1, 7, 1, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-03-01", 2, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-03-01", 3, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-03-01", 4, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-03-01", 5, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-03-01", 6, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-03-01", 7, 7, 2, 2, 0, 0, false)
	TestCalcStartEndTimeStr("2020-02-02", "2021-03-01", 7, 7, 3, 2, 0, 0, false)

	TestCalcStartEndTimeStr("2020-02-02", "2021-02-15", 1, 5, 1, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-02-15", 2, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-02-15", 3, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-02-15", 4, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-02", "2021-02-15", 5, 5, 2, 3, 0, 0, false)
	TestCalcStartEndTimeStr("2020-02-02", "2021-02-15", 5, 5, 3, 3, 0, 0, false)

	fmt.Println("---------------------------------------------")

	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 1, 13, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 2, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 3, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 4, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 5, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 6, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 7, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 8, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 9, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 10, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 11, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 12, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 13, 13, 2, 1, 0, 0, false)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 13, 13, 3, 1, 0, 0, false)

	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 1, 7, 1, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 2, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 3, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 4, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 5, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 6, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 7, 7, 2, 2, 0, 0, false)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 7, 7, 3, 2, 0, 0, false)

	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 1, 5, 1, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 2, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 3, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 4, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 5, 5, 2, 3, 0, 0, false)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 5, 5, 3, 3, 0, 0, false)

	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 1, 4, 1, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 2, 4, 2, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 3, 4, 2, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 4, 4, 2, 4, 0, 0, false)
	TestCalcStartEndTimeStr("2020-01-28", "2021-02-15", 4, 4, 3, 4, 0, 0, false)

	fmt.Println("---------------------------------------------")

	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 1, 13, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 2, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 3, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 4, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 5, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 6, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 7, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 8, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 9, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 10, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 11, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 12, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 13, 13, 2, 1, 0, 0, false)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 13, 13, 3, 1, 0, 0, false)

	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 1, 7, 1, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 2, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 3, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 4, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 5, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 6, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 7, 7, 2, 2, 0, 0, false)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 7, 7, 3, 2, 0, 0, false)

	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 1, 5, 1, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 2, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 3, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 4, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 5, 5, 2, 3, 0, 0, false)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 5, 5, 3, 3, 0, 0, false)

	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 1, 4, 1, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 2, 4, 2, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 3, 4, 2, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 4, 4, 2, 4, 0, 0, false)
	TestCalcStartEndTimeStr("2020-01-29", "2021-02-15", 4, 4, 3, 4, 0, 0, false)

	fmt.Println("---------------------------------------------")

	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 1, 13, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 2, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 3, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 4, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 5, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 6, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 7, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 8, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 9, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 10, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 11, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 12, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 13, 13, 2, 1, 0, 0, false)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 13, 13, 3, 1, 0, 0, false)

	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 1, 7, 1, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 2, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 3, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 4, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 5, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 6, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 7, 7, 2, 2, 0, 0, false)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 7, 7, 3, 2, 0, 0, false)

	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 1, 5, 1, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 2, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 3, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 4, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 5, 5, 2, 3, 0, 0, false)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 5, 5, 3, 3, 0, 0, false)

	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 1, 4, 1, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 2, 4, 2, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 3, 4, 2, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 4, 4, 2, 4, 0, 0, false)
	TestCalcStartEndTimeStr("2020-01-30", "2021-02-15", 4, 4, 3, 4, 0, 0, false)

	fmt.Println("---------------------------------------------")

	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 1, 13, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 2, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 3, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 4, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 5, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 6, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 7, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 8, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 9, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 10, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 11, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 12, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 13, 13, 2, 1, 0, 0, false)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 13, 13, 3, 1, 0, 0, false)

	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 1, 7, 1, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 2, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 3, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 4, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 5, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 6, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 7, 7, 2, 2, 0, 0, false)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 7, 7, 3, 2, 0, 0, false)

	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 1, 5, 1, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 2, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 3, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 4, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 5, 5, 2, 3, 0, 0, false)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 5, 5, 3, 3, 0, 0, false)

	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 1, 4, 1, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 2, 4, 2, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 3, 4, 2, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 4, 4, 2, 4, 0, 0, false)
	TestCalcStartEndTimeStr("2020-01-31", "2021-02-15", 4, 4, 3, 4, 0, 0, false)

	fmt.Println("---------------------------------------------")

	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 1, 13, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 2, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 3, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 4, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 5, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 6, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 7, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 8, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 9, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 10, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 11, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 12, 13, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 13, 13, 2, 1, 0, 0, false)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 13, 13, 3, 1, 0, 0, false)

	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 1, 7, 1, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 2, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 3, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 4, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 5, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 6, 7, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 7, 7, 2, 2, 0, 0, false)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 7, 7, 3, 2, 0, 0, false)

	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 1, 5, 1, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 2, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 3, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 4, 5, 2, 3, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 5, 5, 2, 3, 0, 0, false)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 5, 5, 3, 3, 0, 0, false)

	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 1, 4, 1, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 2, 4, 2, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 3, 4, 2, 4, 0, 0, true)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 4, 4, 2, 4, 0, 0, false)
	TestCalcStartEndTimeStr("2020-02-01", "2021-02-15", 4, 4, 3, 4, 0, 0, false)

	fmt.Println("---------------------------------------------")
	fmt.Println("---------------------------------------------")

	TestCalcStartEndTimeStr("2020-04-01", "2020-06-30", 1, 2, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2020-04-01", "2020-06-30", 2, 2, 2, 1, 0, 0, true)

	TestCalcStartEndTimeStr("2017-12-28", "2018-04-30", 1, 4, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-12-28", "2018-04-30", 2, 4, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-12-28", "2018-04-30", 3, 4, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-12-28", "2018-04-30", 4, 4, 2, 1, 0, 0, true)

	TestCalcStartEndTimeStr("2018-01-30", "2018-04-29", 1, 3, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2018-01-30", "2018-04-29", 2, 3, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2018-01-30", "2018-04-29", 3, 3, 2, 1, 0, 0, false)

	TestCalcStartEndTimeStr("2018-04-10", "2019-05-09", 1, 4, 1, 1, 1, 1, true)
	TestCalcStartEndTimeStr("2018-04-10", "2019-05-09", 2, 4, 2, 1, 1, 1, true)
	TestCalcStartEndTimeStr("2018-04-10", "2019-05-09", 3, 4, 2, 1, 1, 1, true)
	TestCalcStartEndTimeStr("2018-04-10", "2019-05-09", 4, 4, 2, 1, 1, 1, true)

	TestCalcStartEndTimeStr("2018-01-30", "2019-01-29", 1, 12, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2018-01-30", "2019-01-29", 2, 12, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2018-01-30", "2019-01-29", 3, 12, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2018-01-30", "2019-01-29", 4, 12, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2018-01-30", "2019-01-29", 5, 12, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2018-01-30", "2019-01-29", 6, 12, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2018-01-30", "2019-01-29", 7, 12, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2018-01-30", "2019-01-29", 8, 12, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2018-01-30", "2019-01-29", 9, 12, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2018-01-30", "2019-01-29", 10, 12, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2018-01-30", "2019-01-29", 11, 12, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2018-01-30", "2019-01-29", 12, 12, 2, 1, 0, 0, false)

	TestCalcStartEndTimeStr("2017-06-30", "2018-06-29", 1, 6, 1, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-06-30", "2018-06-29", 2, 6, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-06-30", "2018-06-29", 3, 6, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-06-30", "2018-06-29", 4, 6, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-06-30", "2018-06-29", 5, 6, 2, 2, 0, 0, true)
	TestCalcStartEndTimeStr("2017-06-30", "2018-06-29", 6, 6, 2, 2, 0, 0, false)

	TestCalcStartEndTimeStr("2017-07-31", "2018-07-30", 1, 12, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-07-31", "2018-07-30", 2, 12, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-07-31", "2018-07-30", 3, 12, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-07-31", "2018-07-30", 4, 12, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-07-31", "2018-07-30", 5, 12, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-07-31", "2018-07-30", 6, 12, 2, 1, 0, 0, true)

	TestCalcStartEndTimeStr("2017-01-31", "2018-01-30", 1, 12, 1, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-01-30", 2, 12, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-01-30", 3, 12, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-01-30", 4, 12, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-01-30", 5, 12, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-31", "2018-01-30", 6, 12, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-01-30", 7, 12, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-01-30", 8, 12, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-01-30", 9, 12, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-01-30", 10, 12, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-01-30", 11, 12, 2, 1, 0, 0, true)
	TestCalcStartEndTimeStr("2017-01-30", "2018-01-30", 12, 12, 2, 1, 0, 0, false)
}
