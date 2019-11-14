package main

import (
	"fmt"
	"time"
)

//全局变量，创建一个channel
var ch =make(chan int)
var ch1 =make(chan string)

func Printer(str string) {
	for _,value:=range str{
		fmt.Printf("%c",value)
		time.Sleep(time.Microsecond)
	}
}

//person1执行完才执行person2
func person1() {
	Printer("hello")
	ch<-666//给管道发送数据
}

func person2() {
	<-ch//从管道接收数据，如果没有数据会阻塞
	Printer("world")
	ch1<-"子协程结束"//给管道发送数据
}

func main() {
	defer fmt.Println("主协程结束")

	go person1()
	go person2()

	flag:=<-ch1
	fmt.Println(flag)

}
