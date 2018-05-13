package main

import (
	"fmt"
	"strconv"
	"strings"
)

func TestSpilt(billCode string) {
	i, _ := strconv.Atoi(strings.Split(billCode, "-HIS-")[1:][0])
	fmt.Println("i:", i)
}
