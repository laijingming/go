package client

import (
	"config"
	"crawler/engine"
	"crawler_distributed/rpcsupport"
	"crawler_distributed/worker"
	"fmt"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		return nil, err
	}

	return func(r engine.Request) (engine.ParseResult, error) {
		var result worker.ParseResult
		err = client.Call("CrawlerServer.Process", worker.SerializeRequest(r), &result)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeSerializeParseResult(result), nil
	}, nil
}
