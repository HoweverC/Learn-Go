package main

import "fmt"

//定义一个结构体
type Student struct {
	id int
	name string
	sex byte
	age int
	addr string
}

func testStruct1(s Student){
	s.id = 666
	fmt.Println("testStruct1 = ",s)
}

func testStruct2(p *Student){
	p.id = 666
	fmt.Println("testStruct2 = ",p)
}

func main() {
	//顺序初始化，每个成员必须初始化
	var s1 Student = Student{1,"mike",'m',18,"China"}
	fmt.Println("s1 = ",s1)

	//指定成员初始化，没有初始化的成员自动赋值为0
	s2 := Student{name:"mike",addr:"China"}
	fmt.Println("s2 = ",s2)

	//结构体指针变量初始化
	var s3 *Student = &Student{1,"mike",'m',18,"China"}
	fmt.Println("s3 = ",s3)

	s4 := &Student{name:"mike",addr:"China"}
	fmt.Println("s4 = ",s4)

	//操作成员使用.运算符
	var s5 Student
	s5.id = 1
	s5.name ="mike"
	s5.sex ='m'
	s5.age = 18
	s5.addr = "China"
	fmt.Println("s5 = ",s5)

	//指针变量操作成员
	//指针有合法指向后才操作成员
	var s6 Student
	var p1 *Student
	p1 = &s6
	//第一种方法
	p1.id = 2
	//第二种方法，不常用
	(*p1).name = "mike"
	p1.sex = 'm'
	p1.age = 20
	p1.addr = "China"
	fmt.Println("s6 = ",s6 ,"p1 = ",p1)

	//通过new()申请一个结构体
	p2 := new(Student)
	p2.id = 2
	p2.name = "mike"
	p2.sex = 'm'
	p2.age = 20
	p2.addr = "China"
	fmt.Println("p2 = ",p2)

	//结构体比较和赋值
	s7 := Student{1,"mike",'m',18,"China"}
	s8 := Student{1,"mike",'m',18,"China"}
	s9 := Student{2,"mike",'m',18,"China"}
	fmt.Println("s7 == s8 ",s7 == s8)
	fmt.Println("s7 == s9 ",s7 == s9)

	//同类型的两个结构体可以相互赋值
	var s10 Student
	s10 = s7
	fmt.Println("s7 =  ",s7 ,"s10 = ",s10)

	//结构体做函数参数
	s := Student{2,"mike",'m',18,"China"}
	testStruct1(s)
	testStruct2(&s)
}
