package main

import (
	_ "golang.org/x/text/encoding/simplifiedchinese"
	"study/crawler/engine"
	"study/crawler/zhenai"
)

func main() {
	engine.Run(engine.Request{
		"https://www.zhenai.com/zhenghun",
		zhenai.ParserCityList,
	})

	//bytes, err := ioutil.ReadFile("crawler/zhenai/user.html")
	//if err!=nil {
	//	panic(err)
	//}
	//zhenai.ParserProfile(bytes)
}
