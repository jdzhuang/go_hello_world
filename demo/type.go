package demo

import (
	"fmt"
)

type TypeDemo struct {
	Demo
}

func NewTypeDemo() *TypeDemo {
	return &TypeDemo{ Demo{"type", "mysterious type and interface."} }
}

func (d TypeDemo) Do() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)
	s, ok := i.(string)
	fmt.Println(s,ok)
	f, ok := i.(float64)
	fmt.Println(f, ok)
	//following code would panic
	//f = i.(float64)
	//fmt.Println(f)
	
}
