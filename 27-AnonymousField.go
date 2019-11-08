package main

import "fmt"

type  Mystr string//自定义类型

type Person1 struct {
	name string
	sex byte
	age int
}

type Person2 struct {
	name string
	sex byte
	age int
}

type Stu struct {
	Person1//只有类型没有名字，匿名字段，继承了Person里的成员
	id int
	addr string
	name string//同名字段
	int//其他类型匿名字段，基础类型匿名字段
	Mystr
	*Person2//结构体指针类型匿名字段，具体操作跟结构体指针操作相同
}

func main() {
	var s1 Stu =Stu{Person1{
		name: "mike",
		sex:  'm',
		age:  18,
	},1,"China","ben",100,"自定义类型匿名字段",&Person2{name: "mike2", sex:  'm', age:20}}
	//%+v 显示详细
	fmt.Println("s1 = ",s1)

	//匿名字段成员操作与结构体操作相同，使用.运算符
	fmt.Println("s1.name = ",s1.name," s1.id = ",s1.id," s1.Person = ",s1.Person1)

	//同名字段，就近原则
	//可以显式调用
	fmt.Println("s1.Person1.name = ",s1.Person1.name)
	fmt.Println("s1.Person2.name = ",s1.Person2.name)
}
