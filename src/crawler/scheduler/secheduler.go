package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	in chan engine.Request
}

func (s *SimpleScheduler) ReturnChan() chan engine.Request {
	return s.in
}

func (s *SimpleScheduler) SetChan(r chan engine.Request) {
	s.in = r
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	//这里如果不加goroutine，会导致程序阻塞，
	//发送消息给chan时持续阻塞到数据被接收
	go func() {
		s.in <- r
	}()
}
