package engine

import (
	"fmt"
	"log"
	"sync"
)

func MultiRun(seeds ...Request) {
	//seeds初始种子
	in := make(chan Request)
	out := make(chan ParseResult)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ { //开启工作数量
		wg.Add(1)
		go func() {
			for r := range in {
				result, err := worker(r)
				if err != nil {
					continue
				}
				out <- result
			}
			wg.Done()
		}()
	}
	for _, seed := range seeds {
		in <- seed
	}

	go func() {
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
				log.Printf("Got item %v", item)
			}
		}
	}()
	wg.Wait()
	fmt.Println("done")
}
