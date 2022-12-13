package engine

import "crawler/model"

type Parser interface {
	Parser(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}
type Request struct {
	Url string
	//ParserFun func([]byte) ParseResult
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	Items    []model.User
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
