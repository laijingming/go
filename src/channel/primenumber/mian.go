package main

import (
	"fmt"
	"sync"
	"time"
)

//region 方法1

//是否素数
func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func compute(start int, end int, c chan int, wg *sync.WaitGroup) {
	for i := start; i <= end; i++ {
		if i > 1 && isPrime(i) {
			c <- i
		}
	}
	wg.Done()
}

func computePrime1(num int, workerNum int) {
	c := make(chan int)
	var wg sync.WaitGroup
	avgNum := num / workerNum //每次执行数量
	wg.Add(workerNum)
	primeNum := 0
	go func() {
		for _ = range c {
			//fmt.Println("num:", num, "value:", s)
			primeNum++
		}
	}()
	for i := 0; i < workerNum; i++ {
		end := (i + 1) * avgNum
		if end > num {
			end = num
		}
		go compute(i*avgNum, end, c, &wg)
	}
	wg.Wait()
	fmt.Println(primeNum)
	fmt.Println(primeNum)
}

//endregion

//region 方法三
func computePrime3(num int, workerNum int) {
	c := make(chan int)
	var wg sync.WaitGroup
	primeNum := 0
	wg.Add(workerNum)
	for i := 0; i < workerNum; i++ {
		go func() {
			for s := range c {
				if isPrime(s) {
					//fmt.Println("num:", primeNum, " value:", s)
					primeNum++
				}
			}
			wg.Done()
		}()

	}
	go func(wg *sync.WaitGroup) {
		for i := 2; i <= num; i++ {
			c <- i
		}
		close(c)
	}(&wg)
	wg.Wait()
	fmt.Println(primeNum)
}

//endregion

func main() {
	num := 180000
	t1 := time.Now()
	computePrime3(num, 4)
	fmt.Println("computePrime3执行时间:", time.Since(t1))

	t1 = time.Now()
	computePrime2(num)
	fmt.Println("computePrime2执行时间:", time.Since(t1))

	t1 = time.Now()
	computePrime1(num, 4)
	fmt.Println("computePrime1执行时间:", time.Since(t1))
}

//region 方法二

func computePrime2(end int) {
	num := 0
	for i := 2; i <= end; i++ {
		if i > 1 && isPrime(i) {
			//fmt.Println("num:", num, "value:", i)
			num++
		}
	}
	fmt.Println(num)
}

//endregion
