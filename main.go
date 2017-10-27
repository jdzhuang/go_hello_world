package main

import (
	"fmt"
	"runtime"
)

var gDemos []IDemo

func init() {
	gDemos = []IDemo{
		PointerDemo{},
		SliceDemo{},
		VoidDemo{},
	}
}

func onPanic() {
	if r := recover(); r != nil {
		if _, ok := r.(runtime.Error); ok {
			panic(r)
		}
		if s, ok := r.(string); ok {
			fmt.Printf("Logic Err:%s\n", s)
		}
	}
}

func main() {
	defer onPanic()

	for i, v := range gDemos {
		fmt.Printf("==demo[%d]=> %s\n", i, v)
		v.Do()

	}
}
