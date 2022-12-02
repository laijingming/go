package main

import (
	"context"
	"fmt"
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
	var es elasticStruct
	u := `{"name":"wunder", "age": 1}`
	es.initClient()
	es.create("aj407", "1", u)

}

type elasticStruct struct {
	client *elastic.Client
}
type user struct {
	name string
	age  int
}

func (es *elasticStruct) create(index string, id string, doc interface{}) {

	res, err := es.client.Index().
		Index(index).  // 索引名称
		Id(id).        // 指定文档id
		BodyJson(doc). // 可序列化JSON
		Do(context.Background())

	if err != nil {
		panic(err)
	}
	fmt.Println(res)

}
func (es *elasticStruct) initClient() {
	var err error
	es.client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://localhost:9200"))
	if err != nil {
		panic(err)
	}
}
