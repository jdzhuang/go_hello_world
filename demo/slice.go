package demo

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
	//NOTE: array
	numbers := [6]int{1,2,3,4,5,6}
	fmt.Println("Number-Array", numbers)

	//NOTE: tool function(s)
	printSlice := func(s []int, desc string){
		fmt.Printf("cap:%d, len:%d, %v #%s\n", cap(s), len(s), s, desc)
	}
	printSlices := func(s [][]int, desc string){
		for i,v := range s {
			printSlice(v, fmt.Sprintf("row[%d]",i) )
		}
	}

	//NOTE: nil slice
	var slc []int
	printSlice(slc, "nil")

	//NOTE: slice index
	slc = numbers[1:4] // idx: [low, high) 1,2,3
	printSlice(slc, "numbers[1:4]")
	slc = numbers[2:5] // idx: [low, high) 2,3,4
	printSlice(slc, "numbers[2:5]")
	slc = numbers[0:4] // idx: [low, high) 0,1,2,3
	printSlice(slc, "numbers[0:4]")
	slc = numbers[3:6] // idx: [low, high) 3,4,5
	printSlice(slc, "numbers[3:6]")

	//NOTE: slice defaults
	printSlice(numbers[0:], "numbers[0:]") // equivalent to [0, 6)
	printSlice(numbers[:6], "numbers[:6]") // equivalent to [0, 6)
	printSlice(numbers[0:6], "numbers[0:6]")// equivalent to [0, 6)

	//NOTE: 'make' a slice
	slc = make([]int, 3, 5 )
	printSlice(slc, "make([]int, 0, 5)" )
	slc = make([]int, 5 )
	printSlice(slc, "make([]int, 5)" )  //equivalent to 'make([]int, 5, 5)'

	//NOTE: slices of slices
	var slcs [][]int = [][]int{
		[]int{0,1,0,1},
		[]int{0,0,0,1},
		[]int{1,0,1,0},
		[]int{1,0,0,0},
	}
	printSlices(slcs, "slices of slices")

	//NOTE: append()
	slc = append( numbers[:], 99 )
	printSlice( slc, "append( numbers, 99 )" )
	slc = append( numbers[:], 96, 97, 98, 99 )
	printSlice( slc, "append( numbers, 96, 97, 98, 99 )" )

	//NOTE: slice range
	slc = make([]int, 8)
	for i := range slc {
		slc[i] = 1 << uint(i)
	}
	printSlice(slc, "range i")

}
