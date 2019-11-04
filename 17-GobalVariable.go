package main

import "fmt"

//定义在函数外部的变量是全局变量
var a byte

func main() {
	 //定义在大括号中的变量为局部变量，作用域在大括号中
	 //执行到定义变量那句话才分配空间，离开作用域释放

	 //不同作用域允许定义同名变量
	 //使用变量原则，就近原则
	 var a int
	 fmt.Printf("a tpye is %T\n",a)
}
