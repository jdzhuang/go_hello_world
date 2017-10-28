package main

import (
	"fmt"
)

type SliceDemo struct {
	Demo
}

func NewSliceDemo() *SliceDemo {
	return &SliceDemo{ Demo{"slice", "pointer of array."} }
}

func (d SliceDemo) Do() {
	fmt.Println("slicing...")
}
