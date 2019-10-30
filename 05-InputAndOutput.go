package main

import "fmt"

func main() {
	//var a int
	//var b int
	var (
		name    string
		age     int
		married bool
	)
	//Scan从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符
	//Scanf从标准输入扫描文本，根据format参数指定的格式去读取由空白符分隔的值保存到传递给本函数的参数中
	//Scanln类似Scan，它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置

	//fmt.Scan(&a)
	//fmt.Println(a)

	//fmt.Scanf("%d",&b)
	//fmt.Println(b)

	fmt.Scanln(&name, &age, &married)
	fmt.Printf("name:%s age:%d married:%t \n", name, age, married)
}
