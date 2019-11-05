package main

import "fmt"

//指针默认值nil，没有NULL常量
//&取地址 *通过指针访问目标对象
//不要操作没有合法指向的指针
func main() {
	//每个变量有两重含义，变量的内存和变量的地址
	var a int = 10
	//变量的内存
	fmt.Printf("a = %d\n",a)

	//变量的地址
	fmt.Printf("&a = %v\n",&a)

	//保存变量的地址需要使用指针类型
	//*int 保存int的地址
	//**int 保存*int的地址
	var p *int
	p = &a
	fmt.Printf("p = %v ，&a = %v\n",p,&a)

	*p = 123 //操作的是p指向的内存a
	fmt.Printf("*p = %v ，&a = %v\n",*p,&a)

	//使用new函数
	var pointer *int
	//创建一个没有名字的内存（匿名）
	//动态分配
	//GC
	pointer = new(int)
	fmt.Println("*pointer =", *pointer)
	pointer1 := new(int)
	fmt.Println("*pointer =", *pointer1)
}
