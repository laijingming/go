package main

import (
	"fmt"
	"mock"
	real2 "real"
)

const url = "https://www.baidu.com"

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

type PosterRetriever interface {
	Retriever
	Post(url string, m map[string]string) string
}

func session(pr PosterRetriever) string {
	pr.Post(url, map[string]string{
		"name":     "ljm",
		"language": "go",
	})
	return pr.Get(url)
}

func inspect(r Retriever) {
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println(v)
	case real2.Retriever:
		fmt.Printf("UserAgent:%s\n", v.UserAgent)
	default:
		fmt.Println("default：", v)
		return
	}
}

func main() {
	mockRetriever := &mock.Retriever{Contents: "Mozilla/5.0"}
	inspect(mockRetriever)
	fmt.Println("download：", download(mockRetriever))
	fmt.Println("session：", session(mockRetriever))
	inspect(mockRetriever)
	//var r Retriever
	//r = real2.Retriever{
	//	UserAgent: "Mozilla/5.0",
	//	TimeOut:   time.Minute,
	//}
	//inspect(r)
	//// type assertion
	//realRetriever := r.(real2.Retriever)
	//fmt.Println(realRetriever.TimeOut)
}
