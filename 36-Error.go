package main

import (
	"errors"
	"fmt"
)

//用于不致命地的普通错误

func Div(a,b int)(result int,err error)  {
	err = nil
	if b == 0{
		err = errors.New("分母不能为0")
	}else{
		result = a/b
	}
	return
}

func main() {
	err1 := fmt.Errorf("%s","this is normal err1")
	fmt.Println("err1 = ",err1)

	err2 := errors.New("this is normal err2")
	fmt.Println("err2 = ",err2)

	result,err :=Div(10,0)
	if err !=nil{
		fmt.Println("err = ",err)
	}else{
		fmt.Println("result = ",result)
	}
}
