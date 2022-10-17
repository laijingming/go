package zhenai

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	bytes, err := ioutil.ReadFile("./user.html")
	//bytes, err := fetch.Fetch("https://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	list := ParserCityList(bytes)
	fmt.Println(len(list.Items))
	//const citySize = 470
	//if len(list.Requests) != 470 {
	//	t.Errorf("result should have %d requests;but had %d", citySize, len(list.Requests))
	//}
}

func BenchmarkParserCityList(b *testing.B) {
	bytes, err := ioutil.ReadFile("./city_list_text.html")
	if err != nil {
		panic(err)
	}
	const citySize = 470
	for i := 0; i < b.N; i++ {
		list := ParserCityList(bytes)
		if len(list.Requests) != 470 {
			b.Errorf("result should have %d requests;but had %d", citySize, len(list.Requests))
		}
	}
}
