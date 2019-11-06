package main

//需要导入包
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//生成随机数
	//设置种子，如果种子一样，随机数结果一样
	rand.Seed(time.Now().UnixNano())//以当前系统时间作为种子
	//产生随机数
	for i:=0;i<5;i++{
		fmt.Println("rand = ",rand.Intn(100))//限定在100以内
	}
}
