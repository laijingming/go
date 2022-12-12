package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	rpcdemo "rpc"
)

func main() {
	rpc.Register(rpcdemo.DemoService{})
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	} //监听1234端口
	//{"method":"DemoService.Div","params":[{"A":3,"B":4}],"id":1}
	for {
		accept, err := listen.Accept()
		if err != nil {
			log.Printf("accept error:%v\n", err)
			continue
		}
		//后台处理链接
		go jsonrpc.ServeConn(accept)
	}
}
