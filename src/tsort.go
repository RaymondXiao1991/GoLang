package main

import "fmt"

type Output struct {
	BillCode string `json:"bill_code"` // 账单编号
}

type OutputDetail struct {
	Output
	BillName string `json:"bill_name"` // 账单编号
}

type OutputTotal struct {
	OutputDetail
}

// SortAttribute 账单明细排序
func SortAttribute(array interface{}) error {
	if array == nil {
		return nil
	}

	switch array.(type) {
	case *Output:
		billInfo := (array).(*Output)
		fmt.Println("billInfo:", billInfo)
	case *OutputDetail:
		billInfo := (array).(*OutputDetail)
		fmt.Println("billInfo:", billInfo)
	default:
		return nil
	}
	return nil
}
