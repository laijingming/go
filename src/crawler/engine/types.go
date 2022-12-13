package engine

import "crawler/model"

type Parser interface {
	Parse(contents []byte, url string) ParseResult
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

type NewFunParser struct {
	Func func([]byte, string) ParseResult
	Name string
}

func (n *NewFunParser) Parse(contents []byte, url string) ParseResult {
	return n.Func(contents, url)
}

func (n *NewFunParser) Serialize() (name string, args interface{}) {
	return n.Name, nil
}

func CreateNewFunParser() {

}
