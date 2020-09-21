package lession15

import (
	"fmt"
	"math"
)

/*
 方法

Go 没有类。不过你可以为结构体类型定义方法。

方法就是一类带特殊的 "接收者" 参数的函数。

方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。

在此例中，Abs 方法拥有一个名为 v，类型为 Vertex 的接收者。
*/

type Vertext struct {
	X, Y float64
}

func (v Vertext) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Go() {
	v := Vertext{3, 4}

	fmt.Println(v.Abs()) //5
}

/*
方法即函数

记住：方法只是个带接收者参数的函数。

现在这个 Abs 的写法就是个正常的函数，功能并没有什么变化。
*/

func Abs(v Vertext) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Go1() {
	v := Vertext{3, 4}
	fmt.Println(Abs(v)) //5
}

/*
方法（续）

你也可以为非结构体类型声明方法。

在此例中，我们看到了一个带 Abs2 方法的数值类型 MyFloat。

你只能为在同一包内定义的类型的接收者声明方法，
而不能为其它包内定义的类型（包括 int 之类的内建类型）的接收者声明方法。

（译注：就是接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法。）
*/

type MyFloat float64

func (f MyFloat) Abs2() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func Go2() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs2()) //1.4142135623730951
}

/*
指针接收者

你可以为  指针接收者 声明方法。

这意味着对于某类型 T，接收者的类型可以用 *T 的文法。（此外，T 不能是像 *int 这样的指针。）

例如，这里为 *Vertex 定义了 Scale 方法。

指针接收者的方法可以修改接收者指向的值（就像 Scale 在这做的）。

由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。

试着移除第 16 行 Scale 函数声明中的 *，观察此程序的行为如何变化。

若使用值接收者，那么 Scale 方法会对原始 Vertex 值的副本进行操作。
（对于函数的其它参数也是如此。）Scale 方法必须用指针接受者来更改 main 函数中声明
的 Vertex 的值。
*/
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//去掉 * 输出 5，不去掉输出50
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f

}
func Go3() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
