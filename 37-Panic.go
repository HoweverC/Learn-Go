package main

import "fmt"

//用于不可恢复的错误状态，如数组越界，空指针
//我们不应该调用panic函数来报告普通的错误，而应该只把它当作报告致命错误的一种方式
//一般而言当panic异常发生，程序会中断运行，并立即执行在该gorountine中的defer函数
//随后程序崩溃并输出日志信息，包括panic value和函数调用的堆栈跟踪信息

func test1(){
	fmt.Println("1")
}
func test2(x int){
	fmt.Println("2")
	//显示调用panic函数，使程序中断
	//panic("this is panic")
	var a [3]int
	a[x]=100
}
func test3(){
	fmt.Println("3")
}
func main() {
	test1()
	test2(10)//数组越界
	test3()
}
