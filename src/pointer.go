package main

import "fmt"
import "encoding/json"

type CloseBookingReq struct {
	BookingCode     string `json:"booking_code"`      // 下定单编号
	IsReturnEarnest *bool  `json:"is_return_earnest"` // 是否退定金
}

func pointer() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	fmt.Printf("a的变量类型 = %T\n", a)
	fmt.Printf("b的变量类型 = %T\n", b)
	fmt.Printf("c的变量类型 = %T\n", c)

	ptr = &a
	fmt.Printf("a的值为%d\n", a)
	fmt.Printf("ptr指向的值为%d\n", *ptr)

	var bookingReq CloseBookingReq
	var str = `{"booking_code":"111","is_return_earnest":true}`
	if err := json.Unmarshal([]byte(str), &bookingReq); err != nil {
		fmt.Println("请求参数格式不正确")
	}
	fmt.Println(bookingReq.BookingCode)
	fmt.Println(*bookingReq.IsReturnEarnest)

}
