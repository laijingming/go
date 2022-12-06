package view

import (
	"crawler/model"
	"html/template"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	template := template.Must(template.ParseFiles("template.html"))
	data := model.User{
		Id:   "1982028238",
		Name: "水寒",
		Url:  "https://album.zhenai.com/u/1982028238",
	}
	create, err := os.Create("template_temp.html")
	if err != nil {
		return
	}
	err = template.Execute(create, data)
	if err != nil {
		return
	}
}
