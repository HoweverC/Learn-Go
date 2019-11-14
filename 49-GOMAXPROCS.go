package main

import (
	"fmt"
	"runtime"
)

func main() {
	num:=runtime.GOMAXPROCS(4)//指定运算核数
	fmt.Println(num)
	for{
		go fmt.Print(1)
		fmt.Print(0)
	}
}
