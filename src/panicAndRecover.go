package main

import (
	"errors"
	"fmt"
)

// CheckErr 检查错误
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

// PanicAndRecover
func PanicAndRecover() (billID int64, err error) {
	fmt.Println("the program is begining...")
	
	defer func() {
		check := recover()
		if check != nil {
			fmt.Println("Do Delete Error Bill...")
			billID = -1
			err = errors.New("panic error")
		}
	}()

	f1()

	CheckErr(err)

	return billID, err
}

// Normal
func Normal() (billID int64, err error) {
	fmt.Println("the program is begining...")

	err = f1()

	return billID, err
}

func f1() error {
	err := errors.New("panic")
	panic("the program is panic")
	return err
}
