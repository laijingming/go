package main

import (
	"fmt"
	"time"
)

//func createWorker(id int) chan int {
//func createWorker(id int) <-chan int {//只允从chan获得数据
func createWorker(id int) chan<- int { //只允发数据给chan
	c := make(chan int)
	go worker(id, c)
	return c
}

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("worker %d,received %c\n", id, n)
	}
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Microsecond)
}

func bufferedChannel() {
	//创建时，加了3个缓冲区
	c := make(chan int, 3)
	go worker(1, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	close(c) //结束chan
	time.Sleep(time.Microsecond)
}

func main() {
	c := make(chan int, 10)
	fmt.Printf("c=%v,地址=%p", c, &c)

	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
	for s := range c {
		fmt.Println(s)
	}
	time.Sleep(time.Microsecond)
	//chanDemo2()
	//bufferedChannel()
}

func chanDemo2() {
	var workers [10]work
	for i := 0; i < 10; i++ {
		workers[i] = createWorker2(i)
	}
	for i, work := range workers {
		work.in <- 'a' + i
	}
	for i, work := range workers {
		work.in <- 'a' + i
	}
	//wait for all of them
	for _, work := range workers { //每个worker发了俩边，因此这边需要接收俩次
		<-work.do //done直接输出成功，说明chan.in接收完了
		<-work.do //done直接输出成功，说明chan.in接收完了
	}
}

type work struct {
	in chan int
	do chan bool
}

//创建一个worker
func createWorker2(id int) work {
	w := work{
		in: make(chan int),
		do: make(chan bool),
	}
	go worker2(id, w)
	return w
}

//worker
func worker2(id int, w work) {
	for n := range w.in {
		fmt.Println("worker：", id, "，received：", n)
		go func() { //发消息，没有接收会造成堵塞，因此单独开一个协程
			w.do <- true
		}()
	}
}
