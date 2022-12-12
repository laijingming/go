package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServerRpc(host string, jsonRpc interface{}) error {
	rpc.Register(jsonRpc)
	listen, err := net.Listen("tcp", host)
	if err != nil {
		return err
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
	return nil
}

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}
