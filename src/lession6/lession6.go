package lession6

import "fmt"

/*
常量

常量的声明与变量类似，只不过是使用 const 关键字。

常量可以是字符、字符串、布尔值或数值。

常量不能用 := 语法声明。
*/

const Pi = 3.14

func Lession6() {
	const World = "世界"
	fmt.Println("Hello", World)     //Hello 世界
	fmt.Println("Happy", Pi, "Day") //Happy 3.14 Day

	const Truth = true
	fmt.Println("Go rules?", Truth) //Go rules? true

}

/*

数值常量
数值常量是高精度的 值。

一个未指定类型的常量由上下文来决定其类型。

再尝试一下输出 needInt(Big) 吧。

（int 类型最大可以存储一个 64 位的整数，有时会更小。）

（int 可以存放最大64位的整数，根据平台不同有时会更少。）
*/
const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(y float64) float64 {
	return y * 0.1
}

func Lession6_1() {
	fmt.Println(needInt(Small))   //21
	fmt.Println(needFloat(Small)) //0.2
	fmt.Println(needFloat(Big))   //balabala
}
