package main

import (
	"fmt"
	"time"
)

//用select实现超时

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {
		for{
			select {
			case num:=<-ch:
				fmt.Println(num)
			case <-time.After(3*time.Second):
				fmt.Println("超时")
				quit<-true
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch<-i
	}

	<-quit
	fmt.Println("程序结束")
}
