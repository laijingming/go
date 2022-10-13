package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is ccs@www.gmail.com
My email is ccs1@gmail.com
My email is ccs2@gmail.com
`
func main() {
	//regexp.Compile("")
	compile := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.[a-zA-Z0-9.]+`)
	//findString := compile.FindString(text)
	//findString := compile.FindAllString(text,-1)
	findString := compile.FindAllStringSubmatch(text,-1)
	fmt.Println(findString)
}
