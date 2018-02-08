package main

import (
	"time"
)

// 计算两个时间点之间有几个月零几天
func DateUntil(start, end time.Time) (months, days int) {
	daysIn := time.Date(end.Year(), end.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
	months = (end.Year()-start.Year())*12 + int(end.Month()-start.Month())
	if start.Day() == 1 && end.Day() == daysIn {
		months++
		return
	}

	if end.Day() >= start.Day() {
		days = end.Day() - start.Day() + 1
		return
	}

	if end.Day() == daysIn {
		return
	}

	months--
	days = time.Date(start.Year(), start.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day() - start.Day() + end.Day()
	return
}

// 计算两个时间点之间有几个月零几天
func CalTime(start, end time.Time) (months, days int) {
	//start := time.Unix(startDate, 0)
	//end := time.Unix(endDate, 0)
	daysIn := time.Date(end.Year(), end.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
	months = (end.Year()-start.Year())*12 + int(end.Month()-start.Month())
	if start.Day() == 1 && end.Day() == daysIn {
		months++
		return
	}

	if end.Day()+1 >= start.Day() {
		days = end.Day() - start.Day() + 1
		return
	}

	if end.Day() == daysIn {
		return
	}

	months--
	//days = time.Date(end.Year(), end.Month(), 0, 0, 0, 0, 0, time.UTC).Day() - start.Day() + 1 + end.Day()
	daysIn2 := time.Date(end.Year(), end.Month(), 0, 0, 0, 0, 0, time.UTC).Day()
	if daysIn2 >= start.Day() {
		days = daysIn2 - start.Day() + 1 + end.Day()
	} else {
		days = end.Day()
	}

	return
}

// 求两个时间之间的时间差几个月零几天
func CalcTimeBetweenDates2(startDate int64, endDate int64) (months, days int) {
	start := time.Unix(startDate, 0)
	end := time.Unix(endDate, 0)
	daysIn := time.Date(end.Year(), end.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
	months = (end.Year()-start.Year())*12 + int(end.Month()-start.Month())
	if start.Day() == 1 && end.Day() == daysIn {
		months++
		return months, days
	}

	if end.Day() >= start.Day() {
		days = end.Day() - start.Day() + 1
		return months, days
	}

	if end.Day() == daysIn {
		return months, days
	}

	months--
	days = time.Date(start.Year(), start.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day() - start.Day() + end.Day()

	return months, days
}

func testf(start, end string) {
	t1, _ := time.Parse("2006-01-02", start)
	t2, _ := time.Parse("2006-01-02", end)
	months, days := CalTime(t1, t2)
	println(start, end, "=>", months, days)
}

func TestDateUntil() {
	testf("2017-01-29", "2017-02-27")
	testf("2017-01-30", "2017-02-27")
	testf("2017-01-31", "2017-02-27")
	testf("2017-02-01", "2017-02-27")

	testf("2017-01-28", "2017-02-28")
	testf("2017-01-29", "2017-02-28")
	testf("2017-01-30", "2017-02-28")
	testf("2017-01-31", "2017-02-28")
	testf("2017-02-01", "2017-02-28")

	testf("2020-01-29", "2020-02-28")
	testf("2020-01-30", "2020-02-28")
	testf("2020-01-31", "2020-02-28")
	testf("2020-02-01", "2020-02-28")

	testf("2020-01-28", "2020-02-28")
	testf("2020-01-28", "2020-02-29")
	testf("2020-01-29", "2020-02-29")
	testf("2020-01-30", "2020-02-29")
	testf("2020-01-31", "2020-02-29")
	testf("2020-02-01", "2020-02-29")

	testf("2017-12-31", "2018-02-28")
	testf("2017-01-31", "2017-02-26")
	testf("2017-01-31", "2017-03-01")
	testf("2015-05-31", "2018-03-01")
	testf("2020-01-29", "2021-02-28")
	testf("2017-01-29", "2020-02-29")
	testf("2020-01-30", "2020-02-29")
	testf("2020-01-31", "2020-02-29")
	testf("2020-02-01", "2020-02-29")
	testf("2018-12-31", "2019-02-28")
	testf("2020-01-31", "2020-02-26")
	testf("2020-01-31", "2020-03-01")
	testf("2017-03-31", "2017-04-30")
	testf("2017-04-02", "2017-05-01")

	testf("2017-02-02", "2017-02-02")
	testf("2017-02-02", "2018-02-05")
}
