package main

import "fmt"

//select实现斐波那契数列
func fibonacci(ch chan <- int,quit <- chan bool){
	x,y:=1,1
	for{
		select {
		case ch<-x:
			x,y =y,x+y
		case flag:=<-quit:
			fmt.Println(flag)
			return
		}
	}
}

func main() {
	//select监听channel上的数据流动
	//case语句里必须是一个IO操作
	ch:=make(chan int)
	quit:=make(chan bool)

	//消费者，从channel读取内容
	go func() {
		for i:=0;i<8;i++{
			num:=<-ch
			fmt.Println(num)
		}
		quit<-true
	}()

	//生产者
	fibonacci(ch,quit)

}
