package zhenai

import (
	"crawler/engine"
	"regexp"
)

var profileRe = regexp.MustCompile(`<div[^(class)]+class="m-btn[^"]+"[^>]+>([^<]+)</div>`)
var nameRe = regexp.MustCompile(`<h1[^class]+class="nickName"[^>]+>([^<]+)</h1>`)
var genderRe = regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[^>]+">[^>]+(.{1})士征婚</a>`)

func ParserProfile(contents []byte, name string) engine.ParseResult {
	subMatch := profileRe.FindAllSubmatch(contents, -1)
	var infoTemp []string
	infoTemp = append(infoTemp, name)
	for _, sm := range subMatch {
		//fmt.Println(i,string(sm[1]))
		infoTemp = append(infoTemp, string(sm[1]))
	}
	//nikeName:=nameRe.FindSubmatch(contents)
	//gender:=genderRe.FindSubmatch(contents)
	//info:=model.Profile{}
	//info.Marriage=infoTemp[0]
	//info.Age=infoTemp[1]
	//info.Height=infoTemp[3]
	//info.Weight=infoTemp[4]
	//info.Occupation=infoTemp[7]
	//info.Income=infoTemp[6]
	//info.Education=infoTemp[8]
	//info.House=infoTemp[14]
	//info.Car=infoTemp[15]
	//info.Name=string(nikeName[1])
	//info.Gender=string(gender[1])
	result := engine.ParseResult{}
	result.Items = append(result.Items, infoTemp)
	return result

}
