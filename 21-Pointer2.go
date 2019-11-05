package main

import "fmt"

func swap1(a,b int){
	a,b =b,a
}

func swap2(p1,p2 *int){
	*p1,*p2=*p2,*p1
}
func main() {
	a,b := 10,20
	swap1(a,b)//变量本身传递（值传递）
	fmt.Printf("a = %d b = %d\n",a,b)

	swap2(&a,&b)//变量地址传递
	fmt.Printf("a = %d b = %d\n",a,b)
}
