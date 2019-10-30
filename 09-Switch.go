package main

import "fmt"

func main() {
	//switch不需要break
	//默认包含break，可以写
	//fallthrough代表不跳出switch，无条件执行后面的语句
	//支持一个初始化语句，和变量本身以分号分隔
	//case后可以多个条件
	switch num :=1;num{
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
	case 3,4,5:
		fmt.Println("3")
	default:
		fmt.Println("10")
	}

	//switch可以没有条件
	//case后放条件
	score := 100
	switch{
	case score >90 :
		fmt.Println("大于90")
	}
}
