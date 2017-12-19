package main

import "fmt"

//递归
func fact(n int) int {
	if n == 0 {
		return n
	}
	return n + fact(n-1)
}

//指针
func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

//结构体
type person struct {
	name string
	age  int
}

func main() {
	//递归
	sum := fact(100)
	fmt.Println("sum:", sum)

	i := 1
	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("&zeroptr:", &i)

	//结构体
	per := person{name: "zhangjie", age: 12}
	fmt.Println("per:", per)
	fmt.Println("per_name:", per.name)

	sp := &per
	fmt.Println("s_age:", sp.age)
	sp.age = 22
	fmt.Println("sp_age:", sp.age)
}
