package main

import (
	"bufio"
	"fmt"
	"os"
	"study/functional/fib"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}
func writeFile(filename string) {
	//file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666) //创建文件
	file, err := os.Create(filename) //创建文件
	if err != nil {
		//err = errors.New("自己创建的error")
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
		//fmt.Println("error:", err)
		return
	}
	defer file.Close() //关闭文件
	write := bufio.NewWriter(file)
	defer write.Flush() //写入文件
	f := fib.Fibonacci()
	for i := 0; i < 30; i++ {
		fmt.Fprintln(write, f())
	}
}
func main() {
	writeFile("defer.txt")
}
