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
		fmt.Println("mock.Retriever：", v)
	case real2.Retriever:
		fmt.Printf("real2.Retriever：UserAgent:%s\n", v.UserAgent)
	default:
		fmt.Println("default：", v)
		return
	}
}

type person struct {
	name string
	age  int
}

func (p person) String() string {
	return fmt.Sprintf("name：%s,age：%d", p.name, p.age)
}

func main() {
	p := person{
		name: "张胜男",
		age:  0,
	}
	fmt.Println(p)

	//mockRetriever := &mock.Retriever{Contents: "Mozilla/5.0"}
	//inspect(mockRetriever)
	//fmt.Println("download：", download(mockRetriever))
	//fmt.Println("session：", session(mockRetriever))
	//inspect(mockRetriever)

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
