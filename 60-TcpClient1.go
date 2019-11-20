package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn,err:=net.Dial("tcp","127.0.0.1:8000")
	if err!=nil{
		fmt.Println("error = ",err)
		return
	}

	defer conn.Close()

	go func() {
		//从键盘输入内容
		str:=make([]byte,1024)
		for{
			n,err:=os.Stdin.Read(str)//从键盘读取内容放在str
			if err!=nil{
				fmt.Println("error = ",err)
				return
			}
			conn.Write(str[:n-1])
		}
	}()

	buf:=make([]byte,1024)
	for{
		n,err:=conn.Read(buf)
		if err!=nil{
			fmt.Println("error = ",err)
			return
		}
		fmt.Println("buf = ",string(buf[:n-1]))
	}
}
