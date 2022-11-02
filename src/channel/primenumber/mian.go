package main

import (
	"fmt"
	"sync"
)

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return true
		}
	}
	return false
}

func compute(start int, end int, c chan int, do *sync.WaitGroup) {
	for i := start; i < end; i++ {
		if !isPrime(i) {
			c <- i
			do.Add(1)
		}
	}
}

func write(id int, do *sync.WaitGroup) chan int {
	work := make(chan int)
	go func() {
		for s := range work {
			fmt.Println("id:", id, "-", s)
			do.Done()
		}
	}()
	return work
}

func main() {
	var do sync.WaitGroup
	go compute(2, 2000, write(1, &do), &do)
	go compute(2000, 4000, write(2, &do), &do)
	go compute(4000, 6000, write(3, &do), &do)
	go compute(6000, 8000, write(4, &do), &do)
	do.Wait()
}
