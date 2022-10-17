package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "abs我a我睡觉哦abcdefg"
	str, i := lengthOfNonRepeatingSubStr(s)
	fmt.Println(str, i)
	//lengthOfNonRepeatingSubStr2(s)
	//fmt.Println(lengthOfNonRepeatingSubStr3(s))
}

// 寻找最长不含重复字符的字串
func lengthOfNonRepeatingSubStr(str string) (string, int) {
	strMap := make(map[string]int)
	strTemp := ""
	resStr := ""
	for i, s := range str {
		fmt.Println(string(s))
		sTemp := string(s)
		strTemp += sTemp
		if strMap[sTemp] > -1 {
			resStr = strTemp
			strTemp = ""
			strMap[sTemp] = 1
		}
		strMap[sTemp] = i
	}
	return resStr, len(resStr)
}

func lengthOfNonRepeatingSubStr2(str string) {
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
func lengthOfNonRepeatingSubStr3(str string) int {
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
