package main

import "fmt"

func main() {
	//mapStr :=make(map[string]int)
	mapStr := map[string]string{"France": "巴黎", "Beijing": "北京"}
	//遍历1
	for country := range mapStr {
		fmt.Println(country, "中文名：", mapStr[country])
	}
	//遍历2
	mapInt := map[int]string{1: "巴黎", 2: "北京"}
	for country, v := range mapInt {
		fmt.Println(country, "中文名：", v)
	}

	country, ok := mapStr["Beijing"]
	fmt.Println(country, ok)

	//删除key
	delete(mapStr, "Beijing")
	fmt.Println(mapStr)
}
