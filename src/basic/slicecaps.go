package main

import "fmt"

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

//更新切片
func updateSlice(s1 []int) {
	s1[0] = 100
}
func main() {
	var intArr []int
	printSlice(intArr)
	intArr = append(intArr, 1)
	printSlice(intArr)

	intArr2 := make([]int, len(intArr), cap(intArr)*2)
	copy(intArr2, intArr)
	printSlice(intArr2)

	sliceInt := []int{0, 1, 2, 3, 4}
	s1 := sliceInt[1:2] //截取大于等于1小于2
	updateSlice(s1)     //切片会改变原始数据
	fmt.Println(s1)
	fmt.Println(sliceInt)
	fmt.Println("Extending s1--")
	fmt.Println("s1=", s1)
	s2 := s1[0:2]
	fmt.Println("s2=", s2)
	s2 = append(s2, 1000)
	fmt.Println("append 1000 s2=", s2)
	fmt.Println("sliceInt=", sliceInt)
	fmt.Println("delete elements from slic")
	sliceInt = sliceInt[1:]
	printSlice(sliceInt)
	sliceInt = append(sliceInt[:1], sliceInt[2:]...)
	printSlice(sliceInt)
	sliceInt2 := []int{0, 1, 2, 3, 4}
	font := sliceInt2[1:]
	sliceInt3 := font[:len(font)-1]
	printSlice(sliceInt3)
	fmt.Println(sliceInt3[0:])
	sliceInt4 := sliceInt3[2:4]
	fmt.Println(sliceInt4)

}
