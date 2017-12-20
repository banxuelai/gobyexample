package main

import (
	"fmt"
	"sort"
)

type ByLength []string

func main() {
	//排序
	str := []string{"c", "a", "b"}
	sort.Strings(str)
	fmt.Println(str)

	in := []int{2, 1, 4}
	sort.Ints(in)
	fmt.Println(in)
}
