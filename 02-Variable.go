package main
//导入包后一定要使用
import "fmt"

//测试匿名变量
func test() (a ,b ,c int){
	return 1,2,3
}

func main() {
	//变量声明
	//格式 var 变量名 类型
	//变量声明后一定要使用
	var a int

	//自动推导类型必须先初始化，根据初始化的值进行类型推导
	b := 10

	//Println一段一段进行处理
	fmt.Println("a =",a)

	//%T打印变量类型
	//Printf格式化输出
	fmt.Printf("value is %d\ntype is %T \n" , b,b)

	//多重赋值
	i , j := 10 , 20
	//交换两个变量的值
	i , j = j , i
	fmt.Printf("i = %d\nj = %d \n" , i,j)

	var temp int
	//_(下划线)是匿名变量，不做处理，一般配合函数返回值使用
	temp , _ =i , j
	fmt.Println("temp =",temp)

	//test()返回三个值，只给c赋值，剩下两个不做处理
	var c int
	c ,_ ,_ = test()
	fmt.Println("c =",c)
}