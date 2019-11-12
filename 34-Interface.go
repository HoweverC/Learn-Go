package main

import "fmt"

//定义接口
type Humaner interface {
	//方法，只声明不实现，由别的类型（自定义类型）实现
	sayHi()
}

type Student1 struct {
	name string
}

type Teacher1 struct {
	name string
}

func (temp *Student1)sayHi(){
	fmt.Println("student name = ",temp.name)
}

func (temp *Teacher1)sayHi(){
	fmt.Println("Teacher name = ",temp.name)
}

//定义一个普通函数，函数参数为接口类型
//多态
func WhoSayHi(i Humaner){
	i.sayHi()
}

func main() {
	//定义接口变量
	//只要实现此接口方法的类型，该类型的变量就可以给这个类型（接收者类型）赋值
	var i Humaner

	s1:=&Student1{"mike"}
	i = s1
	i.sayHi()

	t1 := &Teacher1{"ben"}
	i = t1
	i.sayHi()

	//调用同一接口，不同形态，多态
	WhoSayHi(s1)
	WhoSayHi(t1)

	//创建一个切片
	x := make([]Humaner,2)
	x[0]=s1
	x[1]=t1
	for _,value := range x{
		value.sayHi()
	}
}
