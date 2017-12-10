package main

import (
	"fmt"
)

func main() {
	// 1517932800 2018-02-07 00:00:00
	// 1515254400 2018-01-07 00:00:00
	m, d := CalcTimeBetweenDates(1515254400, 1517932800)
	fmt.Println("m,d:", m, d)
	
	// 1485960657 2017-2-1 22:50:57
	// 1512917457 2017-12-10 22:50:57
	m, d = CalcTimeBetweenDates(1485960657, 1512917457)
	fmt.Println("m,d:", m, d)

	// 1485960657 2017-2-1 22:50:57
	// 1514645457 2017-12-30 22:50:57
	m, d = CalcTimeBetweenDates(1485960657, 1514645457)
	fmt.Println("m,d:", m, d)
}  
