package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
)

//转二进制
func convertToBin(n int) string {
	result := ""
	for n > 0 {
		mod := n % 2
		n = n / 2
		result = strconv.Itoa(mod) + result
	}
	return result
}

func sum(n ...int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}

func apply(op func(n ...int) int, n ...int) int {
	p := reflect.ValueOf(op).Pointer()
	fName := runtime.FuncForPC(p).Name()
	fmt.Printf("running function %s:\n", fName)
	return op(n...)
}
