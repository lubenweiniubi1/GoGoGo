package lession3

import "fmt"

/*
函数可以没有参数或接受多个参数。
在本例中，add 接受两个 int 类型的参数。

注意  类型  在  变量名 之后。
*/

func add(x int, y int) int {
	return x + y
}

/*
当连续两个或多个函数的已命名形参   类型相同   时，除最后一个类型以外，其它都可以省略。

在本例中，

x int, y int
被缩写为

x, y int
*/
func add2(x, y int) int {
	return x + y
}

func Lession3() {
	fmt.Println(add(42, 13)) //55
}

func Lession3_1() {
	fmt.Println(add2(41, 13)) //54
}

/*
多值返回
函数可以返回任意数量的返回值。

swap 函数返回了两个字符串
*/
func swap(x, y string) (string, string) {
	return y, x
}

func Lession3_2() {
	fmt.Println(swap("hello", " world")) // world hello
}

/*
命名返回值
Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。

返回值的名称应当具有一定的意义，它可以作为文档使用。

没有参数的 return 语句返回已命名的返回值。也就是 直接 返回。

直接返回语句应当仅用在下面这样的短函数中。在长的函数中它们会影响代码的可读性。
*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func Lession3_3() {
	fmt.Println(split(9)) //4 5
}
