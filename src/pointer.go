package main

import "fmt"

func pointer(){
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

}
