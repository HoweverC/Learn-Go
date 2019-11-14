package main

import (
	"fmt"
)

//只写
func producer(out chan<-int)  {
	for i := 0; i < 10; i++ {
		out<-i
	}
}

//只读
func consumer(in <-chan int)  {
	for num:= range in{
		fmt.Println(num)
	}
}


func main() {
	//创建一个双向channel
	ch:=make(chan int)

	//双向channel可以隐式转换为单向
	//单向无法转换为双向
	//根据括号的位置进行设置

	/*
	//只写
	var writeCh chan <-int = ch
	writeCh<-111

	//只读
	var readCh <-chan int = ch
	<-readCh
	*/

	//生产者
	go producer(ch)//channel传参，引用传递

	//消费者
	consumer(ch)
}
