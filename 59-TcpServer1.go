package main

import (
	"fmt"
	"net"
)

func HandleConn(conn net.Conn){
	//获取客户端网络地址信息

	defer conn.Close()

	addr:=conn.RemoteAddr().String()
	fmt.Println("连接成功")

	//读取数据
	buf:=make([]byte,1024)
	for{
		n,err:=conn.Read(buf)
		if err!=nil{
			fmt.Println("error = ",err)
			return
		}
		str:=string(buf[:n])
		fmt.Println("buf = ",str,",ADDR =",addr)
		if"exit" == str{
			fmt.Println("Exit")
			return
		}
		conn.Write([]byte(string(buf[:n])))
	}
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