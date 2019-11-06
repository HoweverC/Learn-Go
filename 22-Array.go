package main

import "fmt"

//数组作为函数参数时，是值传递
func modify1(a [5]int){
	a[0] = 999
}

func modify2(a *[5]int){
	(*a)[0] = 999
}

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

	d := [3][4]int{
		{1:10},
		{4,5,6},
		{7},
	}
	for i:=0;i<3;i++{
		for j:=0;j<4;j++{
			fmt.Printf("d[%d][%d] = %d\n",i,j,d[i][j])
		}
	}

	//数组支持比较，只支持==和!=,比较是否每个元素都一样
	//两个数组比较，数组类型要相同
	e := [5]int{1,2,3,4,5}
	f := [5]int{1,2,3,4,5}

	//数组也可以赋值
	var g [5]int
	g = e
	fmt.Println(e==f)
	fmt.Println(g)

	//数组作为函数参数时是值传递
	//所以数组越大，速度越慢
	modify1(g)
	fmt.Println(g)

	//可以使用数组指针作为参数，这样就是地址传递
	modify2(&g)
	fmt.Println(g)
}
