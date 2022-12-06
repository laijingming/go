package persist

import (
	"context"
	"crawler/model"
	"fmt"
	"github.com/olivere/elastic"
)

type ElasticStruct struct {
	client *elastic.Client
}

func ItemSaver() chan model.User {
	out := make(chan model.User)
	es := InitElastic()
	go func() {
		for {
			user := <-out
			es.save("user", user.Id, user)
			fmt.Printf("Got save #%s item:%v\n", user.Id, user)

		}
	}()
	return out
}

func (es *ElasticStruct) save(index string, id string, doc interface{}) string {
	res, err := es.client.Index().
		Index(index).  // 索引名称
		Id(id).        // 指定文档id
		BodyJson(doc). // 可序列化JSON
		Do(context.Background())

	if err != nil {
		return ""
	}
	return res.Id
}

func (es *ElasticStruct) get(index string, id string) *elastic.GetResult {
	res, err := es.client.Get().
		Index(index). // 索引名称
		Id(id).       // 指定文档id
		Do(context.Background())

	if err != nil {
		panic(err)
	}
	return res
}

func InitElastic() *ElasticStruct {
	e := ElasticStruct{}
	var err error
	e.client, err = elastic.NewClient(
		elastic.SetSniff(false),
		//默认本地9200
		//elastic.SetURL("http://localhost:9200"),
	)
	if err != nil {
		panic(err)
	}
	return &e
}
