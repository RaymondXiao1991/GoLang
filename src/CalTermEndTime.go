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
		startDateStr = time.Unix(startTime, 0).AddDate(0, paymentMonth*(term-1), 0).Unix()
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

func CalTermEndTimeStr(startTime, endTime time.Time, term, totalTerm, billType, paymentMonth int, isTermNormal bool) (startDateStr, endDateStr time.Time) {
	s, e := CalTermEndTime(startTime.Unix(), endTime.Unix(), term, totalTerm, billType, paymentMonth, isTermNormal)
	return time.Unix(s, 0), time.Unix(e, 0)
}

func TestCase(start, end string, term, totalTerm, billType, paymentMonth int, isTermNormal bool) {
	t1, _ := time.Parse("2006-01-02", start)
	t2, _ := time.Parse("2006-01-02", end)
	s, e := CalTermEndTimeStr(t1, t2, term, totalTerm, billType, paymentMonth, isTermNormal)
	fmt.Println(start, end, "=>", term, s.Format("2006-01-02"), e.Format("2006-01-02"))
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
}
