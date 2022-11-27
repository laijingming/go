package scheduler

import "crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueuedScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

func (q *QueuedScheduler) SetChan(r chan engine.Request) {
	//TODO implement me
	panic("implement me")
}

func (q *QueuedScheduler) WorkerReady(worker chan engine.Request) {
	q.workerChan <- worker
}

func (q *QueuedScheduler) ReturnChan() chan engine.Request {
	//TODO implement me
	panic("implement me")
}

func (q *QueuedScheduler) Run() {
	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var actRequest engine.Request
			var actWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				actRequest = requestQ[0]
				actWorker = workerQ[0]
			}
			select {
			case r := <-q.requestChan:
				requestQ = append(requestQ, r)
			case w := <-q.workerChan:
				workerQ = append(workerQ, w)
			case actWorker <- actRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
