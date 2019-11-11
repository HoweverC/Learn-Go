package main

import "fmt"

//可以给任意自定义类型添加方法
//可以任意命名，接收者（reciver）的类型可以是T或*T。基类型T不能是接口或指针
//不支持重载方法，所以不能定义名字相同参数不同的方法
//接收者类型不同即使名字相同也是不同函数

type MethodStruct struct {
	name string
}

type Name struct {
	name string
}

//接收者为普通变量，非指针，值语义，拷贝
func (temp Name) SetValue (n string){
	temp.name = n
}

//接收者为指针变量，引用传递
func (temp *Name) SetPointer (n string){
	temp.name = n
}

func(temp MethodStruct)PrintStruct(){
	fmt.Println("MethodStruct temp = ",temp)
}

func(temp *MethodStruct)SetName(n string){
	temp.name =n
}

//面向过程
func Add1(a,b int)int{
	return  a + b
}

//面向对象，方法，给某个类型绑定一个函数
//换了一个表现形式
type long int
//temp为接收者，即传递一个参数
func(temp long)Add2(other long)long{
	return temp + other
}

func main() {
	result1 :=Add1(1,1)
	fmt.Println(result1)

	//方法调用格式：变量名.函数（参数）
	var a long = 2
	result2 :=a.Add2(3)
	fmt.Println(result2)

	b := MethodStruct{"MethodStruct name"}
	b.PrintStruct()
	b.SetName("123")
	fmt.Println(b)

	//初始化
	c := Name{"name"}
	fmt.Println(c)

	//值语义
	c.SetValue("name1")
	fmt.Println("SetValue：",c)

	//引用语义
	c.SetPointer("name2")
	fmt.Println("SetPointer：",c)

}
