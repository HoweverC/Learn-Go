package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func InitData(s []int)  {
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<len(s) ;i++  {
		s[i] = rand.Intn(100)
	}
}
func main() {
	//切片并不是数组或者数组指针，它通过内部指针和相关属性引用数组片段，进而实现变长方案
	//slice并不是真正意义上的动态数组，而是一个引用类型
	//slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度
	//[low:high:max] low:下标起点 high:下标终点，左闭右开
	//cap = max - low 容量
	//len = high - low 长度
	//操作某个元素，和数组操作方式一样

	array := []int{1,2,3,0,0}
	slice := array[0:3:5]
	fmt.Println("array = ",array)
	fmt.Println("slice = ",slice,"cap = ",cap(slice),"len = ",len(slice))

	//初始化，使用make
	//格式make(切片类型,len,cap)
	//容量可以不写，容量和长度相同
	slice1 := make([]int,5,10)
	fmt.Println(slice1)

	a :=[]int{1,2,3,4,5,6,7,8,9}
	s1 := a[:]//等价于[0:len(a):cap(a)]，全部截取
	fmt.Println("s1 = ",s1,"cap = ",cap(s1),"len = ",len(s1))

	s2 :=a[3:6:7]
	fmt.Println("s2 = ",s2,"cap = ",cap(s2),"len = ",len(s2))

	s3 :=a[:6]//从0开始取6个元素，容量也是6
	fmt.Println("s3 = ",s3,"cap = ",cap(s3),"len = ",len(s3))

	s4 :=a[3:]//从3开始到结尾
	fmt.Println("s4 = ",s4,"cap = ",cap(s4),"len = ",len(s4))

	//append在原切片末尾添加元素
	s5 :=[]int{}
	fmt.Println("s5 = ",s5,"cap = ",cap(s5),"len = ",len(s5))
	s5 = append(s5,1)
	s5 = append(s5,2)
	s5 = append(s5,3)
	fmt.Println("s5 = ",s5,"cap = ",cap(s5),"len = ",len(s5))

	//超过容量会自动扩容，原容量的二倍
	s6 := make([]int,0,1)
	oldCap := cap(s6)
	for i:=0;i<8;i++{
		s6 = append(s6,i)
		if newCap :=cap(s6);oldCap<newCap {
			fmt.Println("oldCap = ",oldCap,"newCap = ",newCap)
			oldCap = newCap
		}
	}

	s7 :=[]int{1,2}
	s8 :=[]int{6,6,6,6,6,6}

	copy(s8,s7)//copy(目标切片，源切片)，把源切片的元素替换到目的切片对应的位置
	fmt.Println("s8 = ",s8)

	//切片作为函数参数是引用传递
	s9 :=make([]int,10)
	InitData(s9)
	fmt.Println("s9 = ",s9)
	sort.Ints(s9)
	fmt.Println("s9 = ",s9)
}
