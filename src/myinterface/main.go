package main

import (
	"fmt"
	"mock"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("www.baidu.com")
}

func main() {
	r := mock.Retriever{Contents: "test"}
	fmt.Println(download(r))
}
