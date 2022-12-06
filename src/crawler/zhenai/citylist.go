package zhenai

import (
	"crawler/engine"
	"fmt"
	"net/http"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) engine.ParseResult {
	compile, err := regexp.Compile(cityListRe)
	if err != nil {
		panic(err)
	}

	result := engine.ParseResult{}
	subMatch := compile.FindAllSubmatch(contents, -1)
	for _, sm := range subMatch {
		//result.Items = append(result.Items, "City:"+string(sm[2]))
		result.Requests = append(result.Requests,
			engine.Request{Url: string(sm[1]), ParserFun: ParserCity})
		//pageUrls := checkUrl(string(sm[1]))
		//if len(pageUrls) > 0 {
		//	for _, pageUrl := range pageUrls {
		//		result.Requests = append(result.Requests,
		//			engine.Request{Url: pageUrl, ParserFun: ParserCity})
		//	}
		//}
	}
	return result
}

func checkUrl(baseUrl string) []string {
	var urls []string
	i := 1
	for {
		pageUrl := fmt.Sprintf(baseUrl+"/%d", i)
		get, err := http.Get(pageUrl)
		if err != nil {
			panic(err)
		}
		if get.StatusCode != http.StatusOK {
			break
		}
		urls = append(urls, pageUrl)
		i++
	}
	return urls
}
