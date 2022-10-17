package zhenai

import (
	"regexp"
	"study/crawler/engine"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^>]+)</a>`

func ParserCity(contents []byte) engine.ParseResult {
	compile, err := regexp.Compile(cityRe)
	if err != nil {
		panic(err)
	}

	result := engine.ParseResult{}
	subMatch := compile.FindAllSubmatch(contents, -1)
	for _, sm := range subMatch {
		name := string(sm[2])
		result.Items = append(result.Items, "User:"+name)
		result.Requests = append(result.Requests,
			engine.Request{Url: string(sm[1]), ParserFun: func(bytes []byte) engine.ParseResult {
				return ParserProfile(bytes, name)
			}})
	}
	return result
}
