package zhenai

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^>]+)</a>`

var userIdRe = regexp.MustCompile(`http://album.zhenai.com/u/(\d+)`)

func ParserCity(contents []byte) engine.ParseResult {
	compile, err := regexp.Compile(cityRe)
	if err != nil {
		panic(err)
	}
	result := engine.ParseResult{}
	subMatch := compile.FindAllSubmatch(contents, -1)
	for _, sm := range subMatch {
		id := userIdRe.FindSubmatch(sm[1])
		result.Items = append(result.Items, model.User{
			Id:   string(id[1]),
			Name: string(sm[2]),
			Url:  string(sm[1]),
		})
		//result.Requests = append(result.Requests,
		//	engine.Request{Url: string(sm[1]), ParserFun: func(bytes []byte) engine.ParseResult {
		//		return ParserProfile(bytes, name)
		//	}})
	}
	return result
}
