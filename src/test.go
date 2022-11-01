package main

import (
	"bufio"
	"fmt"
	"io"
	"mock"
	"os"
	"real"
	"unicode/utf8"
)

const mk_url = "http://www.imooc.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get(mk_url)
}

func Post(poster Poster) {
	poster.Post(mk_url,
		map[string]string{
			"name":     "ljm",
			"language": "go",
		})
}

/*组合接口*/
type RetrieverPoster interface {
	Retriever
	Poster
}

/*组合接口调用*/
func session(rp RetrieverPoster) string {
	rp.Post(mk_url, map[string]string{
		"name":     "ljm",
		"language": "go",
	})
	return rp.Get(mk_url)
}

func main() {
	//r := mock.Retriever{"another fake"}
	//fmt.Println(session(&r))
	//printFile("./src/base.txt")
	//	s:=`s"abc"
	//1
	//在jam沙`
	//printFileContents(strings.NewReader(s))
	//fmt.Println(lengthOfNonRepeatingSubStr2("商店sss"));
	//fmt.Println(lengthOfNonRepeatingSubStr3("商店sss"));
}

/*读内容*/
func printFileContents(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

/*获取文件内容*/
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
}

func inspect(r Retriever) {
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println(v)
	case real.Retriever:
		fmt.Printf("UserAgent:%s\n", v.UserAgent)
	default:
		fmt.Println(v)
		return
	}
}

//更新切片
func updateSlice(s1 []int) {
	s1[0] = 100
}

/*计算字符串未出现重复字符最多的次数*/
func lengthOfNonRepeatingSubStr3(str string) int {
	strMap := make(map[rune]int) //记录字符位置
	start := 0                   //未出现重复字符起始位置
	maxLength := 0
	for i, v := range []rune(str) {
		//strMap字符存在说明字符再次出现，需要更新start，
		//获取记录位置strI,当strI大于等于start，
		//start = strI+1
		if strI, ok := strMap[v]; ok && strI >= start {
			start = strI + 1
		}
		//记录最大长度：当前位置i-start+1
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		strMap[v] = i //记录字符位置
	}
	fmt.Println(strMap)
	return maxLength
}
func lengthOfNonRepeatingSubStr2(str string) int {
	strMap := make(map[string]int) //记录字符位置
	start := 0                     //未出现重复字符起始位置
	maxLength := 0
	for i, v := range []rune(str) {
		//strMap字符存在说明字符再次出现，需要更新start，
		//获取记录位置strI,当strI大于等于start，
		//start = strI+1
		if strI, ok := strMap[string(v)]; ok && strI >= start {
			start = strI + 1
		}
		//记录最大长度：当前位置i-start+1
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		strMap[string(v)] = i //记录字符位置
	}
	fmt.Println(strMap)
	return maxLength
}

func lengthOfNonRepeatingSubStr(str string) {
	strMap := make(map[string]int)
	resultMap := make(map[string]int)
	resultStr := ""
	endResultStr := ""
	lenStr := utf8.RuneCountInString(str)
	for i, s := range []rune(str) {
		str = string(s)
		strMap[str]++
		if strMap[str] < 2 {
			resultStr += str
		}
		if strMap[str] > 1 || i == lenStr {
			if len(strMap) > len(resultMap) {
				resultMap = strMap
				endResultStr = resultStr
			}
			strMap = make(map[string]int)
			resultStr = ""
		}
	}
	fmt.Println(endResultStr)
	fmt.Println(len(resultMap))

}

type phone interface {
	call()
}
type Nokia struct{}
type Iphone struct{}

func (phone Nokia) call() {
	fmt.Println("I am Nokia")
}
func (phone Iphone) call() {
	fmt.Println("I am Iphone")
}
func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-2) + Fibonacci(n-1)
}
func Factorial(x int) int {
	if x > 0 {
		result := x * Factorial(x-1)
		return result
	}
	return 1
}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

type Circle struct {
	radius float64
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
