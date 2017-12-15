package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello_World\n")
	fmt.Println("Hello_World2") //输出自带空格

	//值 字符串 整形  浮点 布尔 逻辑运算符
	fmt.Println("go" + " lang")
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("10.0/3.0=", 10.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
