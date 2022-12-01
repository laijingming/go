package persist

import (
	"fmt"
	"github.com/elastic/go-elasticsearch"
	"strings"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	itemNum := 0
	go func() {
		for {
			item := <-out
			fmt.Printf("Got %d item:%v\n", itemNum, item)
			itemNum++
			save(item)
		}
	}()
	return out
}

func save(item interface{}) {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	})
	if err != nil {
		return
	}
	strArr := strings.Split(item.(string), ":")
	if strArr[0] == "User" {
		es.Index
	}
}
