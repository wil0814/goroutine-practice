package main

import "fmt"

func main() {
	ch, quit := make(chan int), make(chan int)
	go func() {
		ch <- 8   //新增資料
		quit <- 1 //完成
	}()
	for isQuit := false; !isQuit; {
		select {
		case v := <-ch:
			fmt.Printf("%d 來自 ch", v)
		case <-quit:
			isQuit = true
		}
	}
}
