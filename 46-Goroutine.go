package main

import (
	"fmt"
	"time"
)

func newTask() {
	for{
		fmt.Println("this is new Task")
		time.Sleep(time.Second)
	}
}

//主协程推出退出，其它协程也退出
func main() {
	//新协程，使用go关键字
	go newTask()

	for{
		fmt.Println("this is main Task")
		time.Sleep(time.Second)
	}
}
