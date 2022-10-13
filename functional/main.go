package main

import (
	"bufio"
	"fmt"
	"io"
	"study/functional/fib"
)

/*读内容*/
func printFileContents(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fib.Fibonacci()
	printFileContents(f)
}
