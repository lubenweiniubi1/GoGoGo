package lession7

import "fmt"

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
func Lession7_2() {
}
func Lession7_3() {
}
func Lession7_4() {
}
func Lession7_5() {
}
func Lession7_6() {
}
