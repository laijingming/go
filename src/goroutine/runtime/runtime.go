package main

import (
	"fmt"
	"runtime"
)

func main() {

	go func(s string) {
		for i := 0; i < 100000; i++ {
			fmt.Println(s, i)
		}
	}("world-")
	// 主协程
	for i := 0; i < 1; i++ {
		// 切一下，再次分配任务
		runtime.Gosched() //用于让出CPU时间片，让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。
		fmt.Println("hello")
	}
}
