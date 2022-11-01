package main

import (
	"crawler/engine"
	"crawler/zhenai"
	_ "golang.org/x/text/encoding/simplifiedchinese"
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
