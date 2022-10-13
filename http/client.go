package main

import (
	"fmt"
	http "net/http"
	httputil "net/http/httputil"
)

func main() {
	url := "http://www.imooc.com"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	//request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Mobile Safari/537.36")
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:",req)
			return nil
		},
	}
	resp, err :=client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(s))
}
