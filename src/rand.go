package main

import (
    "fmt"
    "math/rand"
    "time"
)

func rrand() {
    for {
        rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
        vcode := fmt.Sprintf("%v", rnd.Int31n(1000000))
        fmt.Println(vcode)
    }
}  
