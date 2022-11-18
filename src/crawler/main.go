package main

import (
	"crawler/engine"
	"crawler/zhenai"
	"fmt"
	_ "golang.org/x/text/encoding/simplifiedchinese"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumGoroutine())
	return
	engine.SimpleEngine{}.Run(engine.Request{
		"https://www.zhenai.com/zhenghun",
		zhenai.ParserCityList,
	})
	//engine.MultiRun(engine.Request{
	//	"https://www.zhenai.com/zhenghun",
	//	zhenai.ParserCityList,
	//})

	//bytes, err := ioutil.ReadFile("crawler/zhenai/user.html")
	//if err!=nil {
	//	panic(err)
	//}
	//zhenai.ParserProfile(bytes)
}
