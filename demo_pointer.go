package main

import (
	"fmt"
)

type PointerDemo struct {
}

func (d PointerDemo) String() string {
	return "Pointer"
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
