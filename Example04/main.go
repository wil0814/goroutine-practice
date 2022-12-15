package main

import (
	"fmt"
	"time"
)

func foo(i int) chan int {
	ch := make(chan int)
	go func() {
		ch <- i
	}()
	return ch
}
func main() {
	ch1, ch2, ch3 := foo(111), foo(222), foo(333)
	ch := make(chan int)
	go func() {
		timeout := time.After(1 * time.Second)
		for isTimeout := false; !isTimeout; {
			select {
			case v1 := <-ch1:
				ch <- v1
			case v2 := <-ch2:
				ch <- v2
			case v3 := <-ch3:
				ch <- v3
			case <-timeout:
				isTimeout = true
			}
		}
	}()
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
}
