package main

import "fmt"

func main() {
	//for 初始化条件;判断条件;条件变化
	//for后面不写条件为无限循环
	sum := 0

	for i := 1 ; i <= 100 ; i++{
		sum += i
	}
	fmt.Println(sum)

	str :="abcde"
	//range
	//迭代默认返回两个值：索引和值
	for index,value:=range str{
		fmt.Printf("str[%d] =%c\n",index,value)
	}
	//不写第二个返回值默认丢弃
	for index:=range str{
		fmt.Printf("str[%d] =%c\n",index,str[index])
	}
	//等价于
	for index,_:=range str{
		fmt.Printf("str[%d] =%c\n",index,str[index])
	}
}
