package main

import "fmt"

//变量它能够调用的方法的集合就是方法集

type S1 struct {
	name string
}

func (s S1)SetValue()  {
	fmt.Println("SetValue")
}

func (s *S1)SetPointer()  {
	fmt.Println("SetPointer")
}

func main() {
	//用实例value和pointer调用方法包括匿名方法不受方法集约束，编译器查找全部方法自动转换reciver实参
	p1 := &S1{"name p1"}
	p2 :=S1{"name p2"}
	//内部转换，把p1转为*p1，然后调用
	//（*p1）.SetValue()
	p1.SetValue()

	p1.SetPointer()

	//内部转换，把*p1转为p1，然后调用
	//p1.SetValue()
	(*p1).SetPointer()

	p2.SetValue()

	//内部转换，把p2转为&p2，然后调用
	//(&p2).SetValue()
	p2.SetPointer()

}
