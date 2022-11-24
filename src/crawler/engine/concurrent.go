package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	SetChan(chan Request)
	ReturnChan() chan Request
}

func (e ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.SetChan(in)
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	itemNum := 0
	for {
		result := <-out
		for _, r := range result.Requests {
			e.Scheduler.Submit(r)
		}
		for _, item := range result.Items {
			itemNum++
			fmt.Printf("Got %d item:%v\n", itemNum, item)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
