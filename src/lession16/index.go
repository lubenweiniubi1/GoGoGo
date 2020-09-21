package lession16

import (
	"fmt"
	"math"
)

/*
指针与函数

现在我们要把 Abs 和 Scale 方法重写为函数。

同样，我们先试着移除掉第 16 的 *。
你能看出为什么程序的行为改变了吗？要怎样做才能让该示例顺利通过编译？

（若你不确定，继续往下看。）
*/

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Go() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	// Scale(v, 10) 这里会报错！！！
	fmt.Println(Abs(v)) //50
}

/*
方法与指针重定向

比较前两个程序，你大概会注意到带指针参数的函数必须接受一个指针：
var v Vertex
ScaleFunc(v, 5)  // 编译错误！
ScaleFunc(&v, 5) // OK

而以指针为接收者的方法被调用时，接收者既能为值又能为指针：
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK

对于语句 v.Scale(5)，即便 v 是个值而非指针，带指针接收者的方法也能被直接调用。
也就是说，由于 Scale 方法有一个指针接收者，为方便起见，Go 会将语句 v.Scale(5)
解释为 (&v).Scale(5)。
*/

func (v *Vertex) Scale1(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc1(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Go1() {
	v := Vertex{3, 4}
	v.Scale1(2)
	ScaleFunc1(&v, 10)

	p := &Vertex{4, 3}
	p.Scale1(3)
	ScaleFunc1(p, 8)

	fmt.Println(v, p) //{60 80} &{96 72}
}

/*
方法与指针重定向（续）
同样的事情也发生在相反的方向。

接受一个值作为参数的函数必须接受一个指定类型的值：
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // 编译错误！

而以值为接收者的方法被调用时，接收者既能为值又能为指针：
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
这种情况下，方法调用 p.Abs() 会被解释为 (*p).Abs()。
*/

/*

选择值或指针作为接收者

使用指针接收者的原因有二：

首先，方法能够修改其接收者指向的值。

其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。

在本例中，Scale 和 Abs 接收者的类型为 *Vertex，即便 Abs 并不需要修改其接收者。

通常来说，所有给定类型的方法都应该有值或指针接收者，但并不应该二者混用。（我们会在接下来几页中明白为什么。）

不多比比 ，接收者为指针更好
*/
