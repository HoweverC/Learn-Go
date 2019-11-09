package main

import "fmt"

func main() {
	//常量 在程序运行时不可以被改变
	//声明使用const关键字
	const a int = 10
	fmt.Println("a =", a)

	//常量也可以自动推导类型
	const b= 10
	fmt.Printf("b type s %T\n", b)

	//多个变量或常量的声明和赋值，同样可以自动推导类型
	var (
		c = 10
		d =3.1415926
	)
	fmt.Println("c =", c)
	fmt.Println("d =", d)

	const (
		i = 1
		j  = 1.23456789
	)
	fmt.Println("i =", i)
	fmt.Println("j =", j)

	//iota常量自动生成器，每隔一行，自行累加一
	//同一行的值相同
	//每次const出现都会清零
	//可以只写一个iota
	//可以使用_（下划线）跳过不需要的值
	const(
		e = iota//0
		f = iota//1
		g		//2
		_		//3（跳过）
		k , l = iota , iota		//4
	)
	fmt.Printf("e = %d ,f = %d ,g = %d k = %d l = %d ", e , f, g ,k ,l)
}