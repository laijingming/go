package model

type SearchResult struct {
	Hits  int
	Start int
	Items []interface{}
}
