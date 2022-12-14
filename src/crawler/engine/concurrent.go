package engine

import "crawler/model"

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan model.User
	RequestProcessor Processor
}
type Processor func(Request) (ParseResult, error)
type Scheduler interface {
	Submit(Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(out, e.Scheduler)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	//itemNum := 0
	for {
		result := <-out
		for _, r := range result.Requests {
			e.Scheduler.Submit(r)
		}
		for _, item := range result.Items {
			//itemNum++
			//fmt.Printf("Got %d item:%v\n", itemNum, item)
			go func(item model.User) { e.ItemChan <- item }(item)
		}
	}
}

func (e *ConcurrentEngine) createWorker(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
