package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("worker %d,received %d\n", id, n)
	}
}
func createWorker(id int) chan<- int { //只允发数据给chan
	c := make(chan int)
	go worker(id, c)
	return c
}
func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			//随机1500毫秒
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Microsecond)
			//time.Sleep(1 * time.Second)
			out <- i
			i++
		}
	}()
	return out
}
func main() {
	//c1, c2 := generator(), generator()
	//for  {
	//	select {
	//	case n := <-c1:
	//		fmt.Println("Received from c1:",n)
	//	case n := <-c2:
	//		fmt.Println("Received from c1:",n)
	//
	//	}
	//}
	c1, c2 := generator(), generator()
	var valuesInt []int
	w := createWorker(1)
	n := 0
	stopTime := time.After(10 * time.Second)
	tk := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(valuesInt) > 1 {
			activeWorker = w
			activeValue = valuesInt[0]
		}
		select {
		case n = <-c1:
			valuesInt = append(valuesInt, n)
		case n = <-c2:
			valuesInt = append(valuesInt, n)
		//case <-time.After(time.Microsecond * 800):
		//	fmt.Println("timeout")
		case activeWorker <- activeValue:
			valuesInt = valuesInt[1:]
		case <-tk:
			fmt.Println("valuesInt len=", len(valuesInt))
		case <-stopTime:
			fmt.Println("stop run")
			return //结束

		}
	}
}
