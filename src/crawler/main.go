package main

import (
	"context"
	"github.com/elastic/go-elasticsearch"
	"github.com/olivere/elastic"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	"https://www.zhenai.com/zhenghun",
	//	zhenai.ParserCityList,
	//})
	//e := &engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.QueuedScheduler{},
	//	WorkerCount: 100,
	//	ItemChan:    persist.ItemSaver(),
	//}
	//e.Run(engine.Request{
	//	Url:       "https://www.zhenai.com/zhenghun",
	//	ParserFun: zhenai.ParserCityList,
	//})

	testElastic()
}

func testElastic() {
	initEsClient()
	es.Index()
}

func Upsert(ctx context.Context, index, id string, doc interface{}) {

}

var esOli *elastic.Client
var es *elasticsearch.Client

func initEsOliClient() {
	var err error
	esOli, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://localhost:9200"))
	if err != nil {
		panic(err)
	}
}
func initEsClient() {
	var err error
	es, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	})
	if err != nil {
		panic(err)
	}
}
