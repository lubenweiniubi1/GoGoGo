package lession13

import (
	"fmt"
	"math"
)

/*
函数值

函数也是值。它们可以像其它值一样传递。

函数值可以用作函数的参数或返回值。
*/

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func Go() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(float64(5), float64(12))) //13

	fmt.Println(compute(hypot)) //5
	fmt.Println(math.Pow)
}

/*
函数的闭包

Go 函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量。

该函数可以访问并赋予其引用的变量的值，换句话说，该函数被这些变量“绑定”在一起。

例如，函数 adder 返回一个闭包。每个闭包都被绑定在其各自的 sum 变量上。
*/

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func Go1() {
	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i)) //
	}
	//最后输出 45 -90
	//sum 会一直叠加，不会被重置
}

/*
斐波那契数列 闭包 ，打印出来
*/

// func fibonacci() func() int {
// 	back1, back2 := 0, 1 // 预先设定好两个初始值

// 	return func() int {

// 		temp := back1                         //记录（back1）的值
// 		back1, back2 = back2, (back1 + back2) // 重新赋值(这个就是核心代码)
// 		return temp                           //返回temp
// 	}
// }

func fibonacci() func() int {
	prev1, prev2 := 0, 1
	return func() int {
		temp := prev1
		prev1, prev2 = prev2, (prev1 + prev2)
		return temp
	}
}

func Fibonacci() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%v  ", f())
	}

}
