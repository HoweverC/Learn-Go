package main

import (
	"fmt"
	"strings"
)

func main() {
	//判断是否包含字串
	fmt.Println("Contains:")
	fmt.Println(strings.Contains("hello go Contains","hello"))
	fmt.Println()

	//组合
	fmt.Println("Join:")
	s1:=[]string{"abc","zxc","Join"}
	buf:=strings.Join(s1,"@")
	fmt.Println(buf)
	fmt.Println()

	//查找索引
	//不存在返回-1
	fmt.Println("Index:")
	fmt.Println(strings.Index("happy go hello index","index"))
	fmt.Println()

	//重复字符串count次，返回重复后的字符串
	fmt.Println("Repeat:")
	buf =strings.Repeat("Repeat",3)
	fmt.Println()

	//分割
	fmt.Println("Split:")
	buf ="hello@abc@go@Split"
	s2:=strings.Split(buf,"@")
	fmt.Println(s2)
	fmt.Println()

	//去除头尾对应的字符
	fmt.Println("Trim:")
	fmt.Println(strings.Trim("@@@Trim@@@","@"))
	fmt.Println()

	//去掉两头空格
	//返回切片
	fmt.Println("Fields:")
	fmt.Println(strings.Fields("   Fields   "))
	fmt.Println()

	//返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串
	fmt.Println("Replace:")
	fmt.Println(strings.Replace("1213141516171819","121314","Replace",3))
}
