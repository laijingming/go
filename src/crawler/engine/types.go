package engine

import "crawler/model"

type Request struct {
	Url       string
	ParserFun func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []model.User
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
