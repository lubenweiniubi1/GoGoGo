package lession8

import (
	"fmt"
	"math"
)

//if
//Go 的 if 语句与 for 循环类似，表达式外无需小括号 ( ) ，而大括号 { } 则是必须的。
func sqrt(x float64) string {
	fmt.Println("运行递归")
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func Lession8() {
	fmt.Println(sqrt(2), sqrt(-4))
}

/*
if 的简短语句
同 for 一样， if 语句可以在条件表达式前执行一个简单的语句。

该语句声明的变量作用域仅在 if 之内。

（在最后的 return 语句处使用 v 看看。）
*/
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
	// return v //undefined
}
func Lession8_1() {
	fmt.Println(
		pow(3, 2, 10), pow(3, 3, 20),
	) //注意最后的 `,`必须加
	// 9 20
}

/*
if 和 else

在 if 的简短语句中声明的变量同样可以在任何对应的 else 块中使用。

（在 main 的 fmt.Println 调用开始前，两次对 pow 的调用均已执行并返回其各自的结果。）
*/
func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func Lession8_2() {
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)

/*
指针

Go 拥有指针。指针保存了值的内存地址。

类型 *T 是指向 T 类型值的指针。其零值为 nil。相当于门牌号
var p *int

& 操作符会生成一个指向其操作数的指针。 取出门牌号
i := 42
p = &i

* 操作符表示指针指向的底层值。 照着门牌号找狗男女

fmt.Println(*p) // 通过指针 p 读取 i
*p = 21         // 通过指针 p 设置 i
这也就是通常所说的“间接引用”或“重定向”。

与 C 不同，Go 没有指针运算。
*/

func Lession8_3() {
	i := 42
	p := &i //取出i的门牌号

	fmt.Println(p)  //0xc000012090
	fmt.Println(*p) //42 照着门牌号找到42 这个狗男女

	*p = 21
	fmt.Println(*p) //21

}

func Lession8_1() {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p) //42
	*p = 21         //同样的房间号换了另一对狗男女
	fmt.Println(i)  //21
	p = &j
	*p = *p / 37
	fmt.Println(j) //73
}
