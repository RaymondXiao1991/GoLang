package main

import (
	"encoding/json"
	"fmt"
)

type Detail struct {
	Type   int     `json:"type"`   //详情的枚举
	Amount float64 `json:"amount"` //金额
}

func mapPriorities() {
	/*
		priorities := map[int]float64{1: 0.00, 3: 0.00, 2: 0.00, 4: 0.00, 5: 0.00}
		details := map[int]float64{1: 100.00, 3: 200.00, 2: 300.00}
		var newDetail map[int]float64

		for type, money := range priorities {
			for k, v := range details {
				if v > 0.00 {
					priority
				}
			}
		}
	*/
	var details = []byte(`[
            {
                "type": 5,
                "amount": 100
            },
            {
                "type": 2,
                "amount": 20
			},
			{
                "type": 4,
                "amount": 300
            }
		]`)
	priorities := []int{1, 3, 2, 4, 5}
	var newDetails []Detail
	var newNewDetails []*Detail
	err := json.Unmarshal(details, &newDetails)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("%v", newDetails)

	for _, priority := range priorities {
		detail := new(Detail)
		for _, v := range newDetails {
			if priority == v.Type {
				detail.Type = priority
				detail.Amount = v.Amount
				newNewDetails = append(newNewDetails, detail)
			}
		}
	}

	for _, v := range newNewDetails {
		fmt.Println(v)
	}

	if data, err := json.Marshal(&newDetails); err != nil {
		fmt.Println("...%v", data)
	}

}
