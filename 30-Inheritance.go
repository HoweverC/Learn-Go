package main

import "fmt"

type Man struct {
	name string
	age int
}

type YoungMan struct {
	Man//匿名字段
	id int
}

func (temp *Man)PrintInfo(){
	fmt.Println("name = ",temp.name,"age = ",temp.age)
}
func main() {
	m := YoungMan{
		Man: Man{"mike", 18},
		id:  100,
	}
	m.PrintInfo()//方法继承自Man
}
