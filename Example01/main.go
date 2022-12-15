package main

import "fmt"

func Intgenerator() chan int {

	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	return ch
}

func main() {
	generator := Intgenerator()
	for i := 0; i < 100; i++ {
		fmt.Println(<-generator)
	}
}
