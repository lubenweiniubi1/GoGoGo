package lession7

import (
	"fmt"
	"math"
)

/*

for
Go 只有一种循环结构：for 循环。

基本的 for 循环由三部分组成，它们用分号隔开：

初始化语句：在第一次迭代前执行
条件表达式：在每次迭代前求值
后置语句：在每次迭代的结尾执行
初始化语句通常为一句短变量声明，该变量声明仅在 for 语句的作用域中可见。

一旦条件表达式的布尔值为 false，循环迭代就会终止。

注意：和 C、Java、JavaScript 之类的语言不同，Go 的 for 语句后面的三个构成部分外没有小括号，
 大括号 { } 则是必须的
*/

func Lession7() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum) //45
}

//初始化语句和后置语句是可选的。
func Lession7_1() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum) //1024
}

//for 是 Go 中的 “while”
//此时你可以去掉分号，因为 C 的 while 在 Go 中叫做 for。
func Lession7_2() {
	sum := 1
	for sum <= 4048 {
		sum += sum
	}
	fmt.Println(sum)
}

//无限循环
//如果省略循环条件，该循环就不会结束，因此无限循环可以写得很紧凑。
func Lession7_3() {
	for {
		fmt.Println(11111)
	}
}

//if
//Go 的 if 语句与 for 循环类似，表达式外无需小括号 ( ) ，而大括号 { } 则是必须的。
func sqrt(x float64) string {
	fmt.Println("运行递归")
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func Lession7_4() {
	fmt.Println(sqrt(2), sqrt(-4))
}
