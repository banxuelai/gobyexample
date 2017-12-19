package main

import "fmt"

//多返回值
func getVals() (int, int) {
	return 3, 8
}

func main() {
	//slice 切片
	s := make([]string, 3)
	fmt.Println(s)

	towD := make([][]int, 3)
	fmt.Println(towD)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append:", s)

	cp := make([]string, len(s))
	copy(cp, s)
	fmt.Println("cp:", cp)

	l1 := s[2:5]
	fmt.Println("l1:", l1)
	l2 := s[:5]
	fmt.Println("l2:", l2)
	l3 := s[2:]
	fmt.Println("l3:", l3)

	de := []string{"1", "2", "3"}
	fmt.Println("de:", de)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		towD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			towD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", towD)

	//map 关联数组/哈希
	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 2
	fmt.Println("m:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	len := len(m)
	fmt.Println("len:", len)

	a_map := map[string]int{"k1": 11, "k2": 222, "k3": 333}
	fmt.Println("a_map:", a_map)

	//range
	for key, val := range a_map {
		fmt.Printf("%s -> %d\n", key, val)
	}

	for _, v := range a_map {
		fmt.Println("list:", v)
	}

	for k, v := range a_map {
		if v == 222 {
			fmt.Println("kk:", k)
		}
	}

	//多返回值
	a, b := getVals()
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	_, c := getVals()
	fmt.Println("c:", c)

}
