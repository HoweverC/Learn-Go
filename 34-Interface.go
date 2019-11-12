package main

import "fmt"

//定义接口
//子集
type Humaner interface {
	//方法，只声明不实现，由别的类型（自定义类型）实现
	sayHi()
}

//超集
type Humaner1 interface {
	Humaner//匿名字段，继承Humaner接口
	sing(s string)
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

func (temp *Student1)sing(s string){
	fmt.Println("student sing = ",s)
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

	var i1 Humaner1
	s2:=&Student1{"mike"}
	i1 = s2
	i1.sayHi()//继承过来的方法
	i1.sing("song")

	//超集可以转化为子集，子集不能转化为超集
	i = i1
	fmt.Print("超集转化为子集：")
	i.sayHi()

	//空接口 万能类型，可以保存任意类型的值
	//空接口不包含任何方法，所有类型都实现空接口
	//如println（arg...interface）
	var i3 interface{} = 1
	fmt.Println("i3 = ",i3)

}
