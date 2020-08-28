package lession5

/*
Go 的基本类型有

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名
 表示一个 Unicode 码点

float32 float64

complex64 complex128 复数
本例展示了几种类型的变量。 同导入语句一样，变量声明也可以“分组”成一个语法块。

int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，
在 64 位系统上则为 64 位宽。
当你需要一个整数值时应使用 int 类型，除非你有特殊的理由使用固定大小或无符号的整数类型。
*/

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func Lession5() {
	fmt.Printf("Type: %T Value:%v\n", ToBe, ToBe)     //Type: bool Value:false
	fmt.Printf("Type: %T Value:%v\n", MaxInt, MaxInt) //Type: uint64 Value:18446744073709551615
	fmt.Printf("Type: %T Value:%v\n", z, z)           //Type: complex128 Value:(2+3i)
}

/*
零值

没有明确初始值的变量声明会被赋予它们的 零值。

零值是：

数值类型为 0，
布尔类型为 false，
字符串为 ""（空字符串）。
*/

func Lession5_1() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s) //0 0 false ""
}

/*
类型转换
表达式 T(v) 将值 v 转换为类型 T。
一些关于数值的转换：

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
或者，更加简单的形式：

i := 42
f := float64(i)
u := uint(f)

与 C 不同的是，Go 在不同类型的项之间赋值时需要显式转换。
试着移除例子中 float64 或 uint 的转换看看会发生什么
会报错
*/
func Lession5_2() {
	i := 42
	f := float64(i)
	u := uint(f)
	fmt.Println(i, f, u) //42 42 42

}

/*
类型推导
在声明一个变量而不指定其类型时（即使用不带类型的 := 语法或 var = 表达式语法），变量的类型由右值推导得出。

当右值声明了类型时，新变量的类型与其相同：

var i int
j := i // j 也是一个 int
不过当右边包含未指明类型的数值常量时，新变量的类型就可能是 int, float64 或 complex128 了，这取决于常量的精度：

i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
尝试修改示例代码中 v 的初始值，并观察它是如何影响类型的。
*/
func Lession5_3() {
	v := 42
	fmt.Printf("v is of type %T\n", v) //v is of type int
	v1 := 42.0
	fmt.Printf("v is of type %T\n", v1) //v is of type float64
	v2 := 0.867 + 0.5i
	fmt.Printf("v is of type %T\n", v2) //v is of type complex128

}
