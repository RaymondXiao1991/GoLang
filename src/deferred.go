package main

import (
	"time"
	"log"
	"fmt"
)

func double(x int)(result int){
	defer func(){
		fmt.Printf("double(%d) = %d\n",x,result)
	}()

	return x+x
}

func triple(x int) (result int) {
    defer func() { result += x }()
    return double(x)
}

func BigSlowOperation(){
	defer Trace("BigSlowOperation()")()
	
	// ...lots of work...
	_ = double(8)
	
	time.Sleep(1*time.Second)	// sleep one second

	return
}

func Trace(msg string)func(){
	start:=time.Now()
	log.Printf("enter %s",msg)
	return func(){
		log.Printf("exit %s (%s)",msg, time.Since(start))
	}
}