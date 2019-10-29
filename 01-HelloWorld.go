//以包作为管理单位
//每个文件都必须先声明包
//程序必须有qie一个main包
//go build 编译代码，生成可执行文件，运行
//go run 不生成可执行文件，直接运行
package main

import "fmt"

//入口
//主函数
func main()  {
	fmt.Print("Hello Go")
}