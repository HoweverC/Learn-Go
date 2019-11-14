package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		fmt.Println("aaa")

		func(){
			defer fmt.Println("ccc")

			//return//终止此函数，输出aaacccbbb
			runtime.Goexit()//终止所在的协程,输出aaaccc

			fmt.Println("ddd")
		}()
		fmt.Println("bbb")
	}()

	for{
		time.Sleep(time.Microsecond)
	}
}
