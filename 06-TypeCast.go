package main

import "fmt"

func main() {
	//bool不能转换为整型
	//整型不能转换为bool
	//不兼容类型
	var flag bool
	flag = true
	fmt.Println(flag)

	var ch byte
	ch ='a'
	var i int
	//类型转换
	i = int(ch)
	fmt.Println(i)
}
