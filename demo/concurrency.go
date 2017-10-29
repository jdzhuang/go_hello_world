package demo

import (
	"fmt"
	"time"
)

type ConcurrencyDemo struct {
	Demo
}

func NewConcurrencyDemo() *ConcurrencyDemo {
	return &ConcurrencyDemo{ Demo{"concurrency", "goroutine and channel."} }
}

func (d ConcurrencyDemo) Do() {
	//NOTE: goroutine basic
	say := func(s string){
		for i:=0; i<5; i++ {
			//time.Sleep(time.Millisecond)
			fmt.Println(s)
		}
	}
	go say("hello")

	//NOTE: with channel
	sum := func(input []int, result chan int){
		r := 0
		for _, v := range input {
			r = r + v
		}
		result <- r
	}
	slc := []int{1,2,3,4,5,6}
	c := make(chan int, 1)
	go sum(slc[:len(slc)/2], c)
	go sum(slc[len(slc)/2:], c)
	x, y := <-c, <-c
	fmt.Println( x, y )
	
	//NOTE: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.  
	//NOTE: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
	fibonacci := func(n int, c chan int){
		x, y := 0, 1
		for i:=0; i<n; i++ {
			x, y = y, y+x
			c <- x
		}
		close(c)
	}
	c = make(chan int, 10)
	go fibonacci(10, c)
	for v := range c {
		fmt.Println(v)
	}

	//NOTE: select channel
	func(){
		tick := time.Tick( 100*time.Millisecond)
		boom := time.After( 500*time.Millisecond)
		for {
			select {
			case <-tick:
				fmt.Print("tick")
			case <-boom:
				fmt.Println("BOOOOM.")
				return
			default:
				fmt.Print(".")
				time.Sleep(50*time.Millisecond)
			}
		}
	}()
}

