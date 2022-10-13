package main

import (
	"fmt"
	"time"
)

func main() {
	//var a [20]int
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for  {
				fmt.Printf("hello from goroutine %d\n",i)
				//a[i]++
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Minute)
	//fmt.Println(a)
}
