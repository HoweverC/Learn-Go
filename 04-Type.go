package main

import "fmt"

func main() {
	//声明字符类型
	var ch byte
	ch = 97
	//%c以字符形式输出
	//字符使用单引号
	fmt.Printf("%c\n",ch)
	ch = 'a'
	fmt.Printf("%c\n",ch)

	//字符串声明
	//字符串使用双引号
	//字符串隐藏了一个结束符\0
	var str string
	str = "abc"
	fmt.Println(str)

	//声明复数类型
	//使用内建函数real()和imag()输出实部虚部
	var t complex128
	t =2.1 + 3.14i
	fmt.Println("real =", real(t)," imag =",imag(t))
}
