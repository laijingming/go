package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Printf("s1=%v,cap=%d\n", s1, cap(s1))
	fmt.Printf("s2=%v,cap=%d\n", s2, cap(s2))
	s3 := append(s2, 11)
	s4 := append(s3, 12)
	fmt.Printf("s3=%v,cap=%d\n", s3, cap(s3))
	fmt.Printf("s4=%v,cap=%d\n", s4, cap(s4))
	fmt.Printf("arr=%v,cap=%d\n", arr, cap(arr))
}
