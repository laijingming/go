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

func write(id int, c chan int) {
	for {
		for s := range c {
			fmt.Println("id:", id, "-", s)
		}
	}
}

func main() {
	num := 2
	for i := 0; i < 4; i++ {
		work := make(chan int)
		go write(i, work)
		go compute(num, num+2000, work)
		num += 2000
	}
	time.Sleep(time.Second)
}
