package main

import "fmt"

type Woman struct {
	name string
	age int
}

type YoungWoman struct {
	Woman//匿名字段
	id int
}

func (temp *Woman)PrintInfo(){
	fmt.Println("Woman name = ",temp.name,"age = ",temp.age)
}

func (temp *YoungWoman)PrintInfo(){
	fmt.Println("YoungWoman name = ",temp.name,"age = ",temp.age)
}
func main() {
	w := YoungWoman{
		Woman: Woman{"mike", 18},
		id:  100,
	}
	w.PrintInfo()//就近原则，方法重写，先找当前类，没有再找父类
	w.Woman.PrintInfo()//显式调用
}
