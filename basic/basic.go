package main

import "fmt"

func variable() {
	var a int
	var b, c = 3, 4
	var s string
	boolTest := true
	fmt.Printf("%d,%s,%d,%d,%d\n", a, s, b, c, boolTest)
}

//枚举类型
func enums() {
	//const(
	//	JAVA = 0
	//	PHP = 1
	//	GO = 2
	//	JAVASCRIPT = 3
	//)

	//一组常量的自增类型
	//const(
	//	JAVA = iota
	//	_
	//	_
	//	PHP
	//	GO
	//	JAVASCRIPT
	//)

	//一组常量的自增表达式
	const (
		JAVA = string('s' + iota)
		_
		_
		PHP
		GO
		JAVASCRIPT
	)
	fmt.Println(JAVA, PHP, GO, JAVASCRIPT)
}
