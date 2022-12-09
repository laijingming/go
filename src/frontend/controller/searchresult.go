package controller

import (
	"context"
	model2 "crawler/model"
	"frontend/model"
	"frontend/view"
	"github.com/olivere/elastic"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type SearchResultHandle struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandle(template string) SearchResultHandle {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	return SearchResultHandle{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

//localhost:8888?search?q=男 23 已婚&from=20
func (h SearchResultHandle) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//strings.TrimSpace 去除空格
	q := strings.TrimSpace(request.FormValue("q"))
	//from是一个整数
	from, err := strconv.Atoi(request.FormValue("from"))
	if err != nil {
		from = 0
	}
	var page model.SearchResult
	page, err = h.getSearchResult(q, from)
	err = h.view.Render(writer, page)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandle) getSearchResult(q string, from int) (model.SearchResult, error) {
	var page model.SearchResult
	do, err := h.client.Search("user").
		Query(elastic.NewQueryStringQuery(q)).
		From(from).
		Size(15).
		Do(context.Background())
	if err != nil {
		return page, err
	}
	page.Query = q
	page.Hits = int(do.TotalHits())
	page.Start = from
	page.Items = do.Each(reflect.TypeOf(model2.User{}))
	page.PrevFrom = from - len(page.Items)
	page.NextFrom = from + len(page.Items)
	if len(page.Items) == 0 {
		page.NextFrom = 0
	}
	return page, nil
}
