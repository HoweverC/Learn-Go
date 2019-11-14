package main

import (
	"fmt"
)

//如果给定了缓冲区量，通道就是异步的，只要缓冲区还有空间，通信就会无阻塞执行
//channel再发数据引起panic，但可以继续读数据
//对于nil channel，收发都会被阻塞
func main() {
	ch:=make(chan int,3)
	fmt.Printf("len = %d , cap = %d\n",len(ch),cap(ch))

	go func() {
		for i:=0;i<5;i++{
			ch<-i
		}
		close(ch)

	}()

	//使用for-range读取channel，当channel关闭时，for循环会自动退出
	// 无需主动监测channel是否关闭，可以防止读取已经关闭的channel
	// 造成读到数据为通道所存储的数据类型的零值
	for num:=range ch{
		fmt.Println("使用range遍历：",num)
	}

	//已关闭的channel会得到零值，如果不确定channel，需要使用ok进行检测
	/*for {
		if num,ok:=<-ch;ok == true{
			fmt.Println(num)
		}else {
			fmt.Println("管道关闭")
			break
		}
	}*/
}
