package worker

import "crawler/engine"

type CrawlerServer struct {
}

func (CrawlerServer) Process(req Request, res *ParseResult) error {
	request, err := DeSerializeRequest(req)
	if err != nil {
		return err
	}
	worker, err := engine.Worker(request)
	if err != nil {
		return err
	}
	*res = SerializeParseResult(worker)
	return nil
}

func (CrawlerServer) ProcessNo(req engine.Request, res *engine.ParseResult) error {

	return nil
}
