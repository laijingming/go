package main

import (
	"fmt"
	"frontend/controller"
	"net/http"
	"os"
)

func main() {
	getwd, _ := os.Getwd()
	fmt.Println(getwd)
	//根目录下所有数据从/src/frontend/view这个目录来
	http.Handle("/", http.FileServer(http.Dir(getwd+"/view")))

	http.Handle("/search",
		controller.CreateSearchResultHandle(getwd+"/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
