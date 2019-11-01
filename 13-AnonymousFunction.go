package main

import "fmt"

func main() {
	//匿名函数

	//使用方法一
	temp :=100
	f1 := func(){
		fmt.Println("匿名函数调用一：",temp)
	}
	f1()

	//使用方法二，使用函数类型赋值
	type FuncType func()
	var f2 FuncType
	f2 = f1
	f2()

	func(){
		fmt.Println("直接调用无参匿名函数：",temp)
	}()//后接小括号代表调用此匿名函数，即传参

	//闭包以引用方式捕获外部变量
	//在里面修改实际修改外面的变量
	//不关心变量是否超出作用域
	func(){
		temp = 999
		fmt.Println("直接调用无参匿名函数，在内部修改temp的值：",temp)
	}()//代表调用此匿名函数，即传参

	//匿名函数有参无返回值
	func(i,j int){
		fmt.Printf("直接调用有参匿名函数，i = %d j= %d\n",i,j)
	}(200,300)

	//匿名函数有参有返回值
	x,y := func(i,j int)(max,min int){
		if i>j{
			max=i
			min=j
		}else{
			max=j
			min=i
		}
		return
	}(300,400)
	fmt.Printf("直接调用有参匿名函数，x= %d y= %d\n",x,y)
}
