package lession10

import "fmt"

/*
defer

defer 语句会将函数推迟到外层函数返回之后执行。

推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
*/
func Lession10() {
	defer fmt.Println("芜湖，推迟执行咯！")

	fmt.Println("hello")

	//hello
	//芜湖，推迟执行咯！
}

/*
defer 栈
推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。

更多关于 defer 语句的信息，请阅读此博文。
https://blog.go-zh.org/defer-panic-and-recover
*/
func Lession10_1() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
	/*
		counting
		done
		9
		8
		7
		6
		5
		4
		3
		2
		1
		0
	*/
}
