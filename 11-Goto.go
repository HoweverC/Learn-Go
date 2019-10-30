package main

import "fmt"

func main() {
	//goto可以在任何地方使用
	fmt.Println("123")
	//goto是关键字，End是标签，自定义，无条件跳转
	//一定要先使用goto定义标签然后再使用
	goto End
	fmt.Println("456")
	End:
		fmt.Println("789")
}
