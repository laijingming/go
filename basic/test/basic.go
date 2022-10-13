package test

import "math"

func calcTriangle(a int, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}

func calcTriangle2(a int, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}
/*计算字符串未出现重复字符最多的次数*/
func lengthOfNonRepeatingSubStr(str string) int {
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
	return maxLength
}
/*计算字符串未出现重复字符最多的次数*/
func lengthOfNonRepeatingSubStr2(str string) int {
	strMap := make(map[rune]int) //记录字符位置
	start := 0                     //未出现重复字符起始位置
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
	return maxLength
}