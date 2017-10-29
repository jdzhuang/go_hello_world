package main

import (
	"fmt"
	"runtime"
	"strconv"
)

var gDemos []IDemo

func init() {
	gDemos = []IDemo{
		NewDemo(TVoid),
		NewDemo(TPointer),
		NewDemo(TArray),
		NewDemo(TSlice),
		NewDemo(TMap),
		NewDemo(TType),
		NewDemo(TConcurrency),
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

func showOpt() (int,error) {
	fmt.Println("====OPTIONS===")
	for i, v := range gDemos {
		fmt.Printf("[%d] %s\n", i, v)
	}
	fmt.Printf("opt? ")
	var o string
	fmt.Scan(&o)
	return strconv.Atoi(o)
}

func main() {
	defer onPanic()

	for opt,err := showOpt(); err==nil &&opt < len(gDemos); {
		fmt.Println("====", gDemos[opt], "====")
		gDemos[opt].Do()
		opt, err = showOpt(); 
	}

}
