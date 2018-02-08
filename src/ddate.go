package main

import "time"

// 获取几天后(精度只到天的时间)
func GetLaterDate(t time.Time, d int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day()+d, 0, 0, 0, 0, time.Local)
}
