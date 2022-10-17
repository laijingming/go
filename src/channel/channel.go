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
	chanDemo()
	//bufferedChannel()
}
