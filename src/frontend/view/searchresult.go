package view

import (
	"frontend/model"
	"html/template"
	"io"
)

type SearchResultView struct {
	template *template.Template
}

func CreateSearchResultView(filename string) SearchResultView {
	return SearchResultView{
		template.Must(template.ParseFiles(filename)),
	}
}

func (s SearchResultView) Render(wr io.Writer, data model.SearchResult) error {
	return s.template.Execute(wr, data)
}
