package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func TimeOfFuncSpend() {
	threadCount := 100000
	fa := 3.233667

	time1 := TestFn(threadCount, fa, func(fa float64) string {
		return strconv.FormatFloat(fa, 'f', 2, 64)
	})
	log.Printf("FormatFloat 耗时:%.4f ms", time1)

	time2 := TestFn(threadCount, fa, func(fa float64) string {
		return fmt.Sprintf("%.2f", fa)
	})
	log.Printf("Sprintf 耗时:%.4f ms", time2)

}

func TestFn(count int, fa float64, fn func(fa float64) string) float64 {
	t1 := time.Now()

	for i := 0; i < count; i++ {
		fn(fa)
	}

	t2 := time.Now()
	return t2.Sub(t1).Seconds() * 1000
}
