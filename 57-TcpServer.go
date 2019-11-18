package main

import (
	"fmt"
	"net"
)

func main() {
	//监听
	listener,err:= net.Listen("tcp","127.0.0.1:8000")
	if err!=nil{
		fmt.Println("error = ",err)
		return
	}
	defer listener.Close()

	//阻塞等待用户请求
	conn,err:=listener.Accept()
	if err!=nil{
		fmt.Println("error = ",err)
		return
	}
	//接收用户请求
	buf:=make([]byte,1024)
	n,err1:=conn.Read(buf)
	if err1 != nil{
		fmt.Println("error = ",err1)
		return
	}
	fmt.Println("buf = ",string(buf[:n]))

	defer conn.Close()
}
