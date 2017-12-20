package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano() / 1000000)

	fmt.Println(time.Now().Format(time.RFC3339))
	t := time.Now().Unix()
	//t为时间戳
	//格式化为指定格式
	s := time.Unix(t, 0).Format("2006-01-02 15:04:05")
	fmt.Println(s)

	//随机数
	fmt.Println(rand.NewSource(100))
	fmt.Println(rand.Intn(rand.Intn(100)))
	fmt.Println(rand.Intn(100))
}
