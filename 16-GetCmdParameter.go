package main

import (
	"fmt"
	"os"
)

func main() {
	//接收用户传递的参数，以字符串形式
	//以空格分隔
	list := os.Args

	n :=len(list)

	fmt.Println("n =",n)

	for i := 0; i < n; i++ {
		fmt.Printf("list[%d] = %s\n",n,list[i])
	}

	for index ,value := range list{
		fmt.Printf("list[%d] = %s\n",index,value)
	}

}
