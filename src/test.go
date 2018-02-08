package main

import "fmt"

func TestAccuracy() {
	fmt.Println("Hello, World!")

	var waitToPay float64
	var billAmount float64 = 9367.86
	var alreadyPaid float64 = 4367.85
	var financialAmount float64 = 0.00
	waitToPay = (billAmount*1000 - alreadyPaid*1000 - financialAmount*1000) / 1000
	fmt.Println(waitToPay)

	waitToPay = (2.01*1000 - 2.00*1000) / 1000
	fmt.Println(waitToPay)
}
