package main

import "fmt"

//函数名首字母小写为private，大写为public
//参数列表不支持默认参数
//只有一个返回值时可以省略小括号，可以不声明返回值变量

//无参无返回值
func Func1(){
	fmt.Println("Func1:无参无返回值函数")
}

//有参无返回值
//定义函数时定义的参数：形参
//参数传递只能由实参传递给形参
func Func2(a int,b string){
	fmt.Printf("Func2:有参无返回值函数,参数1%d,参数2%s\n",a,b)
}

//可变参数列表
//...type不定参数类型
//只能放在形参中的最后一个
//固定参数一定要传参，不定参数视情况而定
func Func3(args...int){
	fmt.Printf("Func3:不定参数个数%d\n",len(args))
	for index ,value :=range args{
		fmt.Printf("args[%d] = %d\n",index,value)
	}
	Func3Test(args...)//不定参数传递，全部参数传递过去
	Func3Test(args[:2]...)//使用切片划分范围传递[0,2),左闭右开
}

func Func3Test(temp ...int){
	for index ,value :=range temp{
		fmt.Printf("temp[%d] = %d\n",index,value)
	}
}

//无参有返回值
//返回值类型卸载小括号和大括号之间
//可以给返回值起别名
func Func4() (result int){
	result = 4
	//等价于return result
	return
}

//多个返回值
func Func5() (a int,b int,c int){
	a,b,c = 4,5,6
	return
}

//有参有返回
func Func6(a,b int)(max,min int){
	if a > b{
		max = a
		min = b
	}else{
		max = b
		min = a
	}
	return
}

//递归函数
//斐波那契数列
func Func7(n int )(result int){
	if n <= 1{
		result = 1
	}else {
		result = Func7(n-1) + Func7(n-2)
	}
	return
}

func Func8(a,b int) int{
	return  a + b
}

//Func9
func ADD(a,b int) int{
	return  a + b
}

//Func9
func SUB(a,b int) int{
	return  a - b
}

//函数也是数据类型，可以通过Type给一个函数类型起名
type Func8Type func(int,int)int//没有函数名没有大括号（方法体）

//函数也是数据类型，可以通过Type给一个函数类型起名
type Func9Type func(int,int)int//没有函数名没有大括号（方法体）

//回调函数 函数有一个参数是函数类型
//多态，调用同一个接口，实现不同效果
func Func9(a , b int,Func9Test Func9Type) (result int){
	fmt.Println("回调函数")
	result = Func9Test(a,b)//这个函数还没实现,必须先定义再调用
	return
}



func main() {
	Func1()//无参无返回值
	Func2(123,"一二三")//传递的参数：实参
	Func3(3,2,1)//不定参数
	fmt.Println(Func4())//调用函数输出返回值

	a,b,c :=Func5()
	fmt.Println(a,b,c)//多个返回值的函数的调用

	max,min :=Func6(10,20)//有参有返回值
	fmt.Println(max,min)//多个返回值的函数的调用

	//递归函数
	//斐波那契数列
	result := 0
	for i := 0; i <= 10; i++ {
		result = Func7(i)
		fmt.Printf("Func7(%d) is: %d\n", i, result)
	}

	//函数类型的使用
	var func8Test Func8Type
	func8Test = Func8//给函数类型赋值
	fmt.Println(func8Test(10,20))//等价于Func8(10,20)

	//回调函数
	//多态
	//后实现
	func9Result :=Func9(10,100,ADD)
	fmt.Println(func9Result)
	func9Result =Func9(10,100,SUB)
	fmt.Println(func9Result)

}
