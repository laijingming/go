package main

/*
TODO:用同一个管道，同时写入并读取50个整数，主程序需要等全部写和读完才能退出管道
*/
import (
	"fmt"
)

func write(c chan int) {
	for i := 50; i >= 0; i-- {
		c <- i
	}
	close(c)
}

func read(c chan int, d chan bool) {
	for s := range c {
		fmt.Println(s)
	}
	d <- true
	close(d)
}

func main() {
	c := make(chan int)
	done := make(chan bool)
	go read(c, done)
	go write(c)

	<-done
}
