package main

import (
	"fmt"
)

// NumInSlice 替代if
func NumInSlice(list []int, a int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// NumNotInSlice 替代if not
func NumNotInSlice(list []int, a int) bool {
	for _, b := range list {
		if b == a {
			return false
		}
	}
	return true
}

func InSlice(billType int) {
	if NumInSlice([]int{1, 3}, billType) {
		fmt.Println("In!")
	} else {
		fmt.Println("Not In!")
	}
}

func NotInSlice(billType int) {
	if NumNotInSlice([]int{1, 3}, billType) {
		fmt.Println("Not In!")
	} else {
		fmt.Println("In!")
	}
}

func TestInSlice(billType int) {
	InSlice(billType)
}

func TestNotInSlice(billType int) {
	NotInSlice(billType)
}
