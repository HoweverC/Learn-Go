package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf:="abc1azc axc aqc"
	//解析规则,解析正则表达式，如果成功
	reg1:=regexp.MustCompile(`a.c`)
	if reg1 == nil{
		//解析失败返回nil
		fmt.Println("regexp error")
		return
	}
	//根据规则提取信息
	//-1匹配所有
	result1:= reg1.FindAllStringSubmatchIndex(buf,-1)
	fmt.Println("result1 = ",result1)
}
