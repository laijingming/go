package zaw

import (
	"crawler/model"
	"fmt"
	"lib/common"
	"regexp"
)

//链接和城市
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[\w]+)" [\w-]+>([\x{4e00}-\x{9fa5}]+)</a>`

//链接和昵称
const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^>]+)</a>`

//CityListParser 城市列表解析
func CityListParser(contents []byte) {
	compile, err := regexp.Compile(cityListRe)
	if err != nil {
		panic(err)
	}
	subMatch := compile.FindAllSubmatch(contents, -1)
	for _, sm := range subMatch {
		fmt.Println("城市：", string(sm[2]), string(sm[1]))
		//todo:城市搜索
		CityParser(common.HttpGet(string(sm[1])))
	}
	fmt.Println(len(subMatch))
}

//CityParser 城市解析
func CityParser(contents []byte) {
	compile, err := regexp.Compile(cityRe)
	if err != nil {
		panic(err)
	}
	subMatch := compile.FindAllSubmatch(contents, -1)
	for _, sm := range subMatch {
		//todo:用户搜索
		userHtml := common.HttpGet(string(sm[1]))
		if userHtml != nil {
			UserParser(userHtml)
		}
	}
	fmt.Println(len(subMatch))
}

var profileRe = regexp.MustCompile(`<div[^(class)]+class="m-btn[^"]+"[^>]+>([^<]+)</div>`)
var nameRe = regexp.MustCompile(`<h1[^class]+class="nickName"[^>]+>([^<]+)</h1>`)
var genderRe = regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[^>]+">[^>]+(.{1})士征婚</a>`)

//UserParser 用户解析
func UserParser(contents []byte) {
	subMatch := profileRe.FindAllSubmatch(contents, -1)
	var infoTemp []string
	for _, sm := range subMatch {
		infoTemp = append(infoTemp, string(sm[1]))
	}
	nikeName := nameRe.FindSubmatch(contents)
	gender := genderRe.FindSubmatch(contents)
	info := model.Profile{}
	info.Marriage = infoTemp[0]
	info.Age = infoTemp[1]
	info.Height = infoTemp[3]
	info.Weight = infoTemp[4]
	info.Occupation = infoTemp[7]
	info.Income = infoTemp[6]
	info.Education = infoTemp[8]
	//info.House = infoTemp[14]
	//info.Car = infoTemp[15]
	info.Name = string(nikeName[1])
	info.Gender = string(gender[1])
	fmt.Println(info)
}
