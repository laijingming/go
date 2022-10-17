package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	fmt.Println(1)
}
