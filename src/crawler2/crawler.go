package main

import (
	"fmt"
	"lib/common"
)

const url = "https://www.zhenai.com/zhenghun"
const backupFilePath = "crawler2/zaw/city.html"

func main() {
	//contents := common.ReadFile(backupFilePath)
	//if contents == nil {
	//	contents = common.HttpGet(url)
	//	fmt.Println("write " + backupFilePath + " done")
	//	common.WriteFile(backupFilePath, contents)
	//}
	//zaw.CityListParser(contents)

	contents := common.HttpGet(url)
	fmt.Println(string(contents))
}
