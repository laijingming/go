package main

import (
	"fmt"
	"time"
)

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return true
		}
	}
	return false
}

func compute(start int, end int, c chan int) {
	for i := start; i <= end; i++ {
		if !isPrime(i) {
			c <- i
		}
	}
}

func write(id int) chan int {
	work := make(chan int)
	go func() {
		for s := range work {
			fmt.Println("id:", id, "-", s)
		}
	}()
	return work
}

func main() {
	var workers [4]chan int
	num := 2
	for i, _ := range workers {
		go compute(num, num+2000, write(i))
		num += 2000
	}
	time.Sleep(time.Second)
}
