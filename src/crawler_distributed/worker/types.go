package worker

import (
	"crawler/engine"
	"crawler/model"
	"crawler/zhenai"
	"errors"
	logpkg "log"
)

type SerializeParse struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializeParse
}
type ParseResult struct {
	Requests []Request
	Items    []model.User
}

func DeSerializeRequest(req Request) (engine.Request, error) {
	parser, err := deSerializeParser(req.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    req.Url,
		Parser: parser,
	}, nil
}

func DeSerializeParseResult(p ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: p.Items,
	}
	for _, req := range p.Requests {
		request, err := DeSerializeRequest(req)
		if err != nil {
			logpkg.Println(err)
			continue
		}
		result.Requests = append(result.Requests, request)
	}
	return result
}

func deSerializeParser(p SerializeParse) (engine.Parser, error) {
	funcParser := &engine.NewFunParser{
		Name: p.Name,
	}
	switch p.Name {
	case "ParserCity":
		funcParser.Func = zhenai.ParserCity
	case "ParserCityList":
		funcParser.Func = zhenai.ParserCityList
	default:
		return nil, errors.New("unknown parser name")
	}
	return funcParser, nil
}

func SerializeRequest(req engine.Request) Request {
	name, args := req.Parser.Serialize()
	return Request{
		Url: req.Url,
		Parser: SerializeParse{
			name,
			args,
		},
	}
}

func SerializeParseResult(res engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: res.Items,
	}
	for _, req := range res.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}
