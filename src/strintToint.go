package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

type OrderNumber string
type OrderRecord struct {
	Num OrderNumber
}

func TestOrderRecord() {
	orderRecord :=
		OrderRecord{
			Num: "111.01",
		}

	sData := *(*string)(unsafe.Pointer(&orderRecord.Num))
	res, err := strconv.ParseFloat(sData, 32)
	if err != nil {
		res = 0.0
	}
	fmt.Println("orderRecord.Num:%.2f", res)
}
