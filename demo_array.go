package main

import (
	"fmt"
)

type ArrayDemo struct {
	Demo
}

func NewArrayDemo() *ArrayDemo {
	return &ArrayDemo{ Demo{"array", "leng-fixed storage."} }
}

func (d ArrayDemo) Do() {
	var sa [2]string
	sa[0],sa[1]="hello","world"
	fmt.Println("String-Array", sa[0], sa[1])
	fmt.Println("String-Array", sa)
	ia := [4]int{2,3,5,7}
	fmt.Println("Prime-Array", ia)
}
