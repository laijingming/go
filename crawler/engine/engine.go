package engine

import (
	"log"
	"study/crawler/fetch"
)

func Run(seeds ...Request) {
	for len(seeds) > 0 {
		request := seeds[0]
		seeds = seeds[1:]
		log.Printf("Fetching %v", request.Url)
		result, err := worker(request)
		if err != nil {
			continue
		}
		seeds = append(seeds, result.Requests...)
		for _, item := range result.Items {
			log.Printf("Got item %v", item)
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
