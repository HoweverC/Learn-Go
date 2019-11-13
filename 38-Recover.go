package main

import "fmt"

func test4(x int){
	//设置recover
	//recover必须在defer中调用
	defer func() {
		//recover()
		// 可以打印panic错误信息
		if err:=recover();err!=nil{
			fmt.Println(err)
		}
	}()
	//显示调用panic函数，使程序中断
	//panic("this is panic")
	var a [3]int
	a[x]=100
	fmt.Println("4")
}

func test5(){
	fmt.Println("5")
}

func main() {
	test4(10)
	test5()
}
