package engine

import (
	"fmt"
)

func MultiRun(seeds ...Request) {
	//seeds初始种子
	in := make(chan Request)
	out := make(chan ParseResult)
	for i := 0; i < 10; i++ { //开启工作数量
		go createWorker2(in, out)
	}
	for _, seed := range seeds {
		in <- seed
	}

	itemNum := 0

	for {
		result := <-out
		fmt.Println(len(result.Requests))
		//有新地请求加入任务
		for _, r := range result.Requests {
			go func(request Request) {
				in <- request
			}(r)
		}
		for _, item := range result.Items {
			itemNum++
			fmt.Printf("Got %d item:%v\n", itemNum, item)
		}
	}
	fmt.Println("done")
}

func createWorker2(in chan Request, out chan ParseResult) {
	for r := range in {
		result, err := worker(r)
		if err != nil {
			continue
		}
		out <- result
	}
}
