package main

import (
	"crawler/engine"
	"crawler/zhenai"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	"https://www.zhenai.com/zhenghun",
	//	zhenai.ParserCityList,
	//})
	engine.MultiRun(engine.Request{
		"https://www.zhenai.com/zhenghun",
		zhenai.ParserCityList,
	})

	//bytes, err := ioutil.ReadFile("crawler/zhenai/user.html")
	//if err!=nil {
	//	panic(err)
	//}
	//zhenai.ParserProfile(bytes)
}
