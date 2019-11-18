package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个定时器
	//设置时间两秒，两秒后向time通道写内容（当前时间）
	timer := time.NewTimer(2*time.Second)
	timer.Reset(time.Second)//定时器重置
	//timer.Stop()//停止定时器
	fmt.Println("当前时间",time.Now())
	//channel没有时数据阻塞
	t := <-timer.C
	fmt.Println(t)

	//定时两秒，阻塞两秒，两秒后产生一个事件，往channel写内容
	<-time.After(2*time.Second)
	fmt.Println("time.After时间到")
}
