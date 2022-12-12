package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
	rpcdemo "rpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)
	var result float64
	err = client.Call("DemoService.Div", rpcdemo.Args{A: 3, B: 4}, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
