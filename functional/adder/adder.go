package main

import "fmt"

/*adder-闭包例子*/
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

/*正统闭包函数*/
type iAdder func(int2 int) (int, iAdder)
func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	Func := adder2(0)
	for i := 0; i < 11; i++ {
		var s int
		s, Func = Func(i)
		fmt.Printf("0+1+....%d=%d\n", i, s)
	}
}
