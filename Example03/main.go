package main

import (
	"fmt"
	"math/rand"
	"time"
)

func doCompute(x int) int {
	//模擬計算時間
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	return 1 + x
}

func branch(x int) chan int {
	ch := make(chan int)
	go func() {
		ch <- doCompute(x)
	}()

	return ch
}

func Recombination(chs ...chan int) chan int {

	ch := make(chan int)
	// for _, c := range chs {
	// 	go func(c chan int) {
	// 		ch <- <-c
	// 	}(c)
	// }
	go func() {
		for i := 0; i < len(chs); i++ {
			select {
			case v1 := <-chs[i]:
				ch <- v1
			}
		}
	}()
	return ch
}

func main() {
	result := Recombination(branch(10), branch(20), branch(30))

	for i := 0; i < 3; i++ {
		fmt.Println(<-result)
	}
}
