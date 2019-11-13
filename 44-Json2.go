package main

import (
	"encoding/json"
	"fmt"
)

type IT1 struct {
	Company string `json:company`//二次编码，如果时"-"，此字段不会输出到屏幕
	Subject []string`json:subject`
	Price float64`json:"price"`//转化为字符串

}

func main() {
	jsonBuf :=  `{
		"company": "school",
			"price": 666,
			"subject": [
			"Go",
			"Java",
			"Python"
]
}`

	var temp IT1
	//第二个参数要地址传递
	err:=json.Unmarshal([]byte(jsonBuf),&temp)
	if err != nil{
		fmt.Println("error")
		return
	}else {
		fmt.Println(temp)
	}


	m:=make(map[string]interface{},3)
	err=json.Unmarshal([]byte(jsonBuf),&m)
	if err != nil{
		fmt.Println("error")
		return
	}else {
		fmt.Println(m)
	}
	var str string
	//类型断言
	for key,value:=range m{
		switch data := value.(type) {
		case string:
			str = data
			fmt.Printf("map[%s]的值类型为string，value = %s\n",key,str)
		case []string:
			fmt.Printf("map[%s]的值类型为[]string，value = %v\n",key,data)
		case float64:
			fmt.Printf("map[%s]的值类型为float64，value = %f\n",key,data)
		case []interface{}:
			fmt.Printf("map[%s]的值类型为[]interface{}，value = %v\n",key,data)
		}
	}
}
