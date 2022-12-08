package view

import (
	model2 "crawler/model"
	"frontend/model"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	template := CreateSearchResultView("template.html")
	data := model.SearchResult{
		Hits:  1,
		Start: 2,
		Items: []model2.User{
			{
				Id:   "1982028238",
				Name: "水寒",
				Url:  "https://album.zhenai.com/u/1982028238",
			},
		},
	}
	create, err := os.Create("template_temp.html")
	if err != nil {
		return
	}
	err = template.Render(create, data)
	if err != nil {
		return
	}
}
