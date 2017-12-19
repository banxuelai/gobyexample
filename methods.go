package main

import (
	"fmt"
)

type me struct {
	width  int
	height int
}

func (m *me) sum() int {
	return m.height * m.width
}
func main() {
	m := me{width: 2, height: 3}
	fmt.Println(m)
	fmt.Println(m.sum())
}
