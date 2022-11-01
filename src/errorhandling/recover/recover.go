package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println(err.Error())
		} else {
			panic(r)
		}
	}()
	//b := 0
	//a := 5 / b
	//fmt.Println(a)
	panic(123)
}

func main() {
	defer fmt.Println("第一")
	defer fmt.Println("第二")
	tryRecover()
}
