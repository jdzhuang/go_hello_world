package demo

import (
	"fmt"
)

type PointerDemo struct {
	Demo
}

func NewPointerDemo() *PointerDemo {
	return &PointerDemo{ Demo{"pointer", "typical pointer concept."} }
}

func (d PointerDemo) Do() {
	a, b := 12, 48
	p := &a
	fmt.Println("*p", *p)
	*p = a + 6
	fmt.Println("a", a)
	p = &b
	*p = *p / 6
	fmt.Println("b", b)
}
