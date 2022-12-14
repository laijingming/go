package main

import (
	"crawler_distributed/rpcsupport"
	"crawler_distributed/worker"
	"fmt"
	"testing"
	"time"
)

func TestCrawlerServerProcess(t *testing.T) {
	const host = ":1234"
	go Server(host)
	time.Sleep(time.Second * 2)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	var result worker.ParseResult
	err = client.Call("CrawlerServer.Process",
		worker.Request{
			Url: "http://www.zhenai.com/zhenghun/shapingba",
			Parser: worker.SerializeParse{
				Name: "ParserCity",
				Args: nil,
			},
		}, &result)
	if err != nil {
		t.Errorf("err:%s", err)
	} else {
		fmt.Println(result)
	}
}
