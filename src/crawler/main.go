package main

import (
	"context"
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/zhenai"
	"fmt"
	"github.com/olivere/elastic"
)

func main() {

	//engine.SimpleEngine{}.Run(engine.Request{
	//	"https://www.zhenai.com/zhenghun",
	//	zhenai.ParserCityList,
	//})
	e := &engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url: "https://www.zhenai.com/zhenghun",
		Parser: &engine.NewFunParser{
			Func: zhenai.ParserCityList,
			Name: "ParserCityList",
		},
	})

	//e := InitElastic()
	//e.create("aj407", "1", User{
	//	Name: "test",
	//	Id:   "1",
	//})
}

type User struct {
	Name string `json:"name"`
	Id   string `json:"id"`
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

type elasticStruct struct {
	client *elastic.Client
}

var client elasticStruct

func InitElastic() *elasticStruct {
	if client.client != nil {
		return &client
	}
	e := elasticStruct{}
	var err error
	e.client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://localhost:9200"))
	if err != nil {
		panic(err)
	}
	client = e
	return &e
}
