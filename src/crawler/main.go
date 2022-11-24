package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	"https://www.zhenai.com/zhenghun",
	//	zhenai.ParserCityList,
	//})

	engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}.Run(engine.Request{
		Url:       "https://www.zhenai.com/zhenghun",
		ParserFun: zhenai.ParserCityList,
	})

	//bytes, err := ioutil.ReadFile("crawler/zhenai/user.html")
	//if err!=nil {
	//	panic(err)
	//}
	//zhenai.ParserProfile(bytes)
}
