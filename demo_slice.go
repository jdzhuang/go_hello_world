package main

import (
	"fmt"
)

type SliceDemo struct {
}

func (d SliceDemo) String() string {
	return "Slice"
}

func (d SliceDemo) Do() {
	fmt.Println("slicing.")
}
