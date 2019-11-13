package main

import (
	"fmt"
	"strconv"
)

func main() {
	//转换为字符串后追加到字节数组
	slice:=make([]byte,0,1024)
	//第二个参数为要追加的东西，第三个参数为指定以10进制的方式追加
	slice = strconv.AppendInt(slice,123456,10)
	slice = strconv.AppendQuote(slice,"hellogo")
	fmt.Println("slice = ",string(slice))

	//其他类型转换为字符串
	var str string
	str = strconv.FormatBool(false)
	//f指打印格式，小数，-1指小数点位数，64指float664
	str = strconv.FormatFloat(3.14,'f',-1,64)
	fmt.Println(str)

	//整型转字符串
	str = strconv.Itoa(6666)
	fmt.Println(str)

	//字符串转其它类型
	var flag bool
	var err error
	flag,err = strconv.ParseBool("true")
	if err == nil{
		fmt.Println("flag = ",flag)
	}else {
		fmt.Println("error = ",err)
	}

	//字符串转整型
	a,_ := strconv.Atoi("789")
	fmt.Println(a)
}
