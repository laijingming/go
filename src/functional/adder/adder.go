package main

import "fmt"

/*
	adder-闭包例子
	调用adder(),返回func(v int) int,每次调用对sum进行累加存储（sum相当于静态变量）
*/
func adder() func(int) int {
	sum := 0 //自由变量
	return func(v int) int {
		sum += v
		return sum
	}
}

/*正统闭包函数*/
type iAdder func(int) (int, iAdder)

/*
调用adder2(n),返回iAdder
每次调用iAdder(x),返回n+x,adder2(n+x)，此时base=n+x
*/
func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

//斐波那契数列
func fbi() func() int {
	v1, v2 := 0, 1
	return func() int {
		v1, v2 = v2, v1+v2
		fmt.Println(v1)
		return v1
	}
}

func main() {
	Func := adder2(0)
	for i := 0; i < 11; i++ {
		var s int
		s, Func = Func(i)
		fmt.Printf("0+1+....%d=%d\n", i, s)
	}

	Func2 := adder()
	for i := 0; i < 11; i++ {
		fmt.Printf("0+1+....%d=%d\n", i, Func2(i))
	}

	fbiFunc := fbi()
	fbiFunc()
	fbiFunc()
	fbiFunc()
	fbiFunc()
	fbiFunc()
	fbiFunc()
	fbiFunc()

}
