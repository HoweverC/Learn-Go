package main

import "fmt"

func main() {
	//可以先声明
	//if支持一个初始化语句，初始化语句和判断条件用分号隔开
	if a:=10;a == 10{
		fmt.Println("a == 10")
	}else if a > 10{
		fmt.Println("a > 10")
	}else{
		fmt.Println("a < 10")
	}
}
