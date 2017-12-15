package main

import (
	"fmt"
	"math"
	"time"
)

//自定义函数
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
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

	//变量
	var str = "gogorun"
	fmt.Println(str)

	var b, c int = 1, 2
	fmt.Println("b+c=", b+c)

	var testInt int
	fmt.Println(testInt)

	//简洁声明初始化
	test := 123
	fmt.Println(test)

	//常量
	const n = 500000000
	fmt.Println(math.Sin(n))

	x := 1
	y := 2
	fmt.Println(Max(x, y))

	//for
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	//if/else
	if x == 1 && y == 2 {
		fmt.Println("nimei")
	} else if y > 3 {
		fmt.Println(string(y) + "type2")
	} else {
		fmt.Println(string(x) + "type3")
	}
	fmt.Println(time.Now())

	//switch
	sw := 3
	switch sw {
	case 1:
		fmt.Println(sw)
	case 2:
		fmt.Println(string(sw) + "type")
	default:
		fmt.Println(3)
	}

	//array
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 10
	fmt.Println("all:", a)
	fmt.Println("this:", a[4])
	var double_arr [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			double_arr[i][j] = i + j
		}
	}

	fmt.Println("double_arr:", double_arr)

}
