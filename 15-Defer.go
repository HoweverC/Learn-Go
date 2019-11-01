package main

import "fmt"

func main() {
	//defer只能放在函数内部
	//在main函数结束前调用
	fmt.Println("123")
	defer  fmt.Println("789")
	fmt.Println("456")

	//多个defer的执行顺序
	//LIFO后进先出（栈）
	//哪怕函数或者某个延迟调用发生错误，这些调用依旧会执行
	defer fmt.Println("abc")
	defer fmt.Println("def")
	defer fmt.Println("ghi")

	a := 10
	b := 20
	//最后执行，a,b的值已经改变，输出111，222
	defer func(){
		fmt.Printf("外部变量 a=%d,b=%d\n",a,b)
	}()

	//先传参，但不执行，所以a=10,b=20
	defer func(a,b int){
		fmt.Printf("外部变量 a=%d,b=%d\n",a,b)
	}(a,b)

	a = 111
	b = 222
}
