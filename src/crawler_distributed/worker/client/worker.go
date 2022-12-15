package client

import (
	"crawler/engine"
	"crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(rpcClient chan *rpc.Client) (engine.Processor, error) {

	return func(r engine.Request) (engine.ParseResult, error) {
		var result worker.ParseResult
		c := <-rpcClient
		err := c.Call("CrawlerServer.Process", worker.SerializeRequest(r), &result)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeSerializeParseResult(result), nil
	}, nil
}
