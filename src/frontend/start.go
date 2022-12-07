package main

import (
	"frontend/controller"
	"net/http"
)

func main() {
	//根目录下所有数据从/src/frontend/view这个目录来
	http.Handle("/", http.FileServer(http.Dir("./src/frontend/view")))

	http.Handle("/search",
		controller.CreateSearchResultHandle("./src/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
