package engine

import (
	"crawler/fetch"
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	for len(seeds) > 0 {
		request := seeds[0]
		seeds = seeds[1:]
		result, err := worker(request)
		if err != nil {
			continue
		}
		seeds = append(seeds, result.Requests...)
		for _, item := range result.Items {
			log.Printf("Got item %v \n", item)
		}
	}
}

func worker(request Request) (ParseResult, error) {
	log.Printf("Fetching %v", request.Url)
	bytes, err := fetch.Fetch(request.Url)
	if err != nil {
		return ParseResult{}, err
	}
	return request.ParserFun(bytes), nil
}
