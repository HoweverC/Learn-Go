package main

import "fmt"

func main() {
	//数组，同一个类型的集合
	//通过下标操作数组
	//数组定义时，元素个数必须是常量
	var id [5]int
	for i:=0;i<len(id);i++{
		id[i] = i
		fmt.Println(id[i])
	}

	for index,value := range id{
		fmt.Println(index,value)
	}

	//数组初始化，声明同时赋值
	var a [5]int =[5]int{1,2,3,4,5}
	fmt.Println(a)

	//部分初始化
	b := [5]int{1,2,3}
	fmt.Println(b)

	//指定某个元素初始化
	c := [5]int{2:10,4:20}
	fmt.Println(c)


}
