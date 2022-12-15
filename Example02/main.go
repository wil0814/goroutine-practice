package main

import "fmt"

func SendNotifi(user string) chan string {
	//之後可以加資料庫讓他用user去搜資料

	notifi := make(chan string, 500)
	go func() {
		notifi <- fmt.Sprintf("hello %s, welcome", user)
	}()
	return notifi
}
func main() {
	user1 := SendNotifi("william")
	user2 := SendNotifi("zhuang")

	fmt.Println(<-user1)
	fmt.Println(<-user2)
}
