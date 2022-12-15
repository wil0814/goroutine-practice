package main

import "fmt"

func IntGenerator() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func Filter(in chan int, number int) chan int {
	out := make(chan int)
	go func() {
		for {
			i := <-in
			if i%number != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	const max = 3
	numbers := IntGenerator()
	number := <-numbers
	for number <= max {
		fmt.Println(number)
		numbers = Filter(numbers, number)
		number = <-numbers
	}
}
