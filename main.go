package main

import (
	"fmt"
	"runtime"
	"strconv"
	"./demo"
)

var gDemos []demo.IDemo

func init() {
	gDemos = []demo.IDemo{
		demo.NewDemo(demo.TVoid),
		demo.NewDemo(demo.TPointer),
		demo.NewDemo(demo.TArray),
		demo.NewDemo(demo.TSlice),
		demo.NewDemo(demo.TMap),
		demo.NewDemo(demo.TType),
		demo.NewDemo(demo.TConcurrency),
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
