package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai"
	itemsaver "crawler_distributed/persist/client"
	"crawler_distributed/worker/client"
)

func main() {

	//engine.SimpleEngine{}.Run(engine.Request{
	//	"https://www.zhenai.com/zhenghun",
	//	zhenai.ParserCityList,
	//})
	saver, err := itemsaver.ItemSaver()
	if err != nil {
		panic(err)
	}
	processor, err := client.CreateProcessor()
	if err != nil {
		panic(err)
	}
	e := &engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         saver,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url: "https://www.zhenai.com/zhenghun",
		Parser: &engine.NewFunParser{
			Func: zhenai.ParserCityList,
			Name: "ParserCityList",
		},
	})

}
