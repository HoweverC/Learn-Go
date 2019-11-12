package main

import "fmt"

type Student2 struct {
	name string
	id int
}
func main() {
	i := make([]interface{},3)
	i[0] = 100//int
	i[1] = "hello"//string
	i[2] = Student2{"mike",888}
	fmt.Println("使用if实现：")
	for index,value := range i{
		if data,ok :=value.(int); ok == true{
			fmt.Printf("i[%d]的类型为int，%d\n",index,data)
		} else if data ,ok := value.(string); ok == true{
			fmt.Printf("i[%d]的类型为string，%s\n",index,data)
		}else if data ,ok := value.(Student2); ok == true{
			fmt.Printf("i[%d]的类型为struct，%v\n",index,data)
		}
	}
	fmt.Println("使用switch实现：")
	for index,value := range i{
		switch data := value.(type) {
		case int:
			fmt.Printf("i[%d]的类型为int，%d\n",index,data)
		case string:
			fmt.Printf("i[%d]的类型为string，%s\n",index,data)
		case Student2:
			fmt.Printf("i[%d]的类型为struct，%v\n",index,data)
		}
	}
}
