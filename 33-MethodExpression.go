package main

import "fmt"

type Boy struct {
	name string
	age int
}

type YoungBoy struct {
	Boy//匿名字段
	id int
}

func (temp *Boy)PrintInfo(){
	fmt.Println("Boy name = ",temp.name,"age = ",temp.age)
}

func (temp *YoungBoy)PrintInfo(){
	fmt.Println("YoungBoy name = ",temp.name,"age = ",temp.age)
}
func main() {
	y := YoungBoy{
		Boy: Boy{"mike", 18},
		id:  100,
	}
	y.PrintInfo()//传统调用

	//方法表达式
	yFunc :=(*YoungBoy).PrintInfo
	yFunc(&y)//显式把接收者传递过去

}
