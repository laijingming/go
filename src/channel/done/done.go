package main

import (
	"fmt"
	"sync"
)

type woker struct {
	in   chan int
	done func()
}

//func createWorker(id int) chan int {
//func createWorker(id int) <-chan int {//只允从chan获得数据
//func createWorker(id int) chan<- int { //只允发数据给chan
func createWorker(id int, wg *sync.WaitGroup) woker { //只允发数据给chan
	w := woker{
		make(chan int),
		func() {
			wg.Done()
		},
	}
	go doWorker(id, w)
	return w
}

func doWorker(id int, w woker) {
	for n := range w.in {
		fmt.Printf("doWorker %d,received %c\n", id, n)
		w.done()
	}
}

func chanDemo() {
	var w sync.WaitGroup
	var workers [10]woker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &w)
	}
	w.Add(20)
	for i, w := range workers {
		w.in <- 'a' + i
	}

	for i, w := range workers {
		w.in <- 'A' + i
	}
	w.Wait()
}

func main() {
	chanDemo2()
	//bufferedChannel()
}

func chanDemo2() {
	var workers [10]work2
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i] = createWorker2(i, &wg)
	}
	for i, work := range workers {
		work.in <- 'a' + i
	}
	for i, work := range workers {
		work.in <- 'a' + i
	}
	wg.Wait()
}

type work2 struct {
	in   chan int
	done func()
}

//创建一个worker
func createWorker2(id int, wg *sync.WaitGroup) work2 {
	w := work2{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go worker2(id, w)
	return w
}

//worker
func worker2(id int, w work2) {
	for n := range w.in {
		fmt.Println("worker：", id, "，received：", n)
		w.done()
	}
}
