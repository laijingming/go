package fib

import (
	"io"
	"strconv"
	"strings"
)

/*斐波那契数列*/
func Fibonacci() IntGen {
	pre, nex := 0, 1
	return func() int {
		pre, nex = nex, nex+pre
		return pre
	}
}

type IntGen func() int

func (g IntGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 1000000000 {
		return 0, io.EOF
	}
	return strings.NewReader(strconv.Itoa(next) + "\n").Read(p)
}
