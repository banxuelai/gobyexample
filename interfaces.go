package main

import (
	"fmt"
)

//接口
type all interface {
	sum() int
	mul() int
}

type test1 struct {
	x int
	y int
}

type test2 struct {
	t int
}

func (t test1) sum() int {
	return t.x + t.y
}

func (t test1) mul() int {
	return t.x * t.y
}

func (t test2) sum() int {
	return t.t + t.t
}

func (t test2) mul() int {
	return t.t * t.t
}

func mer(a all) {
	fmt.Println(a)
	fmt.Println(a.sum())
	fmt.Println(a.mul())
}

func main() {
	t1 := test1{x: 2, y: 4}
	t2 := test2{t: 5}

	mer(t1)
	mer(t2)
}
