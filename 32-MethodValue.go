package main

import "fmt"

type Girl struct {
	name string
	age int
}

type YoungGirl struct {
	Girl//匿名字段
	id int
}

func (temp *Girl)PrintInfo(){
	fmt.Println("Girl name = ",temp.name,"age = ",temp.age)
}

func (temp *YoungGirl)PrintInfo(){
	fmt.Println("YoungGirl name = ",temp.name,"age = ",temp.age)
}
func main() {
	g := YoungGirl{
		Girl: Girl{"mike", 18},
		id:  100,
	}
	g.PrintInfo()//传统调用
	gFunc1 := g.PrintInfo//这就是方法值，调用函数时无需接收者，隐藏了接收者
	gFunc1()
	g.Girl.PrintInfo()//传统调用
	gFunc2 := g.Girl.PrintInfo
	gFunc2()//方法值调用
}
