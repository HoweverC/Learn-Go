package main

import "fmt"

func main() {
	//类型别名 给类型起一个别名
	//给int64起名bigint
	type bigint int64
	var a bigint
	fmt.Printf("a tpye %T",a)
}
