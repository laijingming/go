package main

import (
	"fmt"
)

func main() {
	a := make(map[int]int)
	for i := 0; i < 10; i++ {
		go func(i int) { //与main并发执行
			for {
				fmt.Printf("hello from goroutine %d\n", i)
				a[i]++
				//runtime.Gosched()
			}
		}(i)
	}
	//time.Sleep(time.Second)//不加sleep可能goroutine还没执行main就已经结束
	fmt.Println(a)
}
