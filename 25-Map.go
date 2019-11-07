package main

import "fmt"

func testMap(m map[int]string){
	delete(m,2)
}

func main() {
	//定义
	//map是无序的键值对，键值唯一
	var m1 map[int]string
	fmt.Println(m1)
	//对于map，只有len没有cap
	fmt.Println(len(m1))

	//通过make创建
	m2:=make(map[int]string)
	fmt.Println(m2)
	fmt.Println(len(m2))

	//使用make可以指定长度,容量可以自动扩容
	m3:=make(map[int]string,10)
	fmt.Println(m3)
	fmt.Println(len(m3))

	//赋值时如果是已经存在的key，修改内容,不存在，追加
	m4:=map[int]string{1:"one",2:"two",3:"three"}
	fmt.Println(m4)
	m4[1]="test"
	fmt.Println(m4)
	fmt.Println(len(m4))

	//遍历
	for key,value :=range m4{
		fmt.Println("key = ",key,"value = ",value)
	}

	//判断key是否存在
	//第一个返回值为key对应的value，第二个返回值为key是否存在，存在ok==true
	value,ok :=m4[1]
	if ok ==true {
		fmt.Println("m4[1] = ",value)
	}else {
		fmt.Println("key不存在")
	}

	//删除
	//删除key为1的内容
	delete(m4,1)

	//map做函数参数
	//引用传递
	testMap(m4)
	fmt.Println(m4)
}
