package zhenai

import (
	"crawler/fetch"
	"fmt"
	"testing"
)

func TestParserCity(t *testing.T) {
	bytes, _ := fetch.Fetch("http://www.zhenai.com/zhenghun/shanxi")
	city := ParserCity(bytes)
	fmt.Printf("%+v", city)
}
