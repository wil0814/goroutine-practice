package main

import (
	"fmt"
	"time"
)

func Timer(SleepTime time.Duration) chan bool {
	ch := make(chan bool)
	go func() {
		time.Sleep(SleepTime)
		ch <- true
	}()
	return ch
}
func main() {
	timeout := Timer(5 * time.Second)
	for {
		select {
		case <-timeout:
			fmt.Println("timeout!!!")
			return
		}
	}
}
