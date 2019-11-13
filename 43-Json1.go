package main

import (
	"encoding/json"
	"fmt"
)

//成员首字母必须大写
type IT struct {
	Company string `json:company`//二次编码，如果时"-"，此字段不会输出到屏幕
	Subject []string
	Price float64`json:",string"`//转化为字符串

}

func main() {
	//通过结构体生成json
	//定义一个结构体变量同时初始化
	s:=IT{"Struct", []string{"GO", "c++", "python"}, 100}

	//编码，根据内容生成json文本

	//普通
	//buf,err :=json.Marshal(s)

	//格式化编码
	buf,err :=json.MarshalIndent(s,"","  ")
	if err != nil{
		fmt.Println("err = ",err)
		return
	}else{
		fmt.Println("buf = ",string(buf))
	}

	m:=make(map[string]interface{},3)
	m["company"] = "Map"
	m["subjects"] = []string{"Go","Java","Python"}
	m["price"] = 666

	//编码
	buf,err =json.MarshalIndent(m,"","  ")
	if err != nil{
		fmt.Println("err = ",err)
		return
	}else{
		fmt.Println("buf = ",string(buf))
	}
}
