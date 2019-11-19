package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn){
	//获取客户端网络地址信息

	defer conn.Close()

	addr:=conn.RemoteAddr().String()
	fmt.Println("连接成功")

	//读取数据
	buf:=make([]byte,2048)
	n,err:=conn.Read(buf)
	if err!=nil{
		fmt.Println("error = ",err)
		return
	}
	fmt.Println("buf = ",strings.Trim(string(buf),""),",ADDR =",addr)
	conn.Write([]byte(string(buf[:n])))
}

func main() {
	//监听
	listener,err:= net.Listen("tcp","127.0.0.1:8000")
	if err!=nil{
		fmt.Println("error = ",err)
		return
	}
	defer listener.Close()

	for{
		conn,err:=listener.Accept()
		if err!=nil{
			fmt.Println("error = ",err)
			return
		}

		//处理用户请求,新建一个协程
		go HandleConn(conn)
	}
}