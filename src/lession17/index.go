package lession17

import (
	"fmt"
	"math"
)

/*
接口

接口类型 是由一组方法签名定义的集合。

接口类型的变量可以保存任何实现了这些方法的值。

注意: 示例代码的 22 行存在一个错误。
由于 Abs 方法只为 *Vertex （指针类型）定义，因此 Vertex（值类型）并未实现 Abser
*/
type Abser interface {
	Abs() float64
}

func Go() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	fmt.Println(a.Abs())

	a = &v
	fmt.Println(a.Abs())

	// a = v 报错
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
接口与隐式实现

类型通过实现一个接口的所有方法来实现该接口。
既然无需专门显式声明，也就没有“implements”关键字。

隐式接口从接口的实现中解耦了定义，这样接口的实现可以出现在任何包中，无需提前准备。

因此，也就无需在每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义。
*/
type I interface {
	M()
}

type T struct {
	S string
}

//此方法表示类型T实现了接口I，但是我们无需显示的声明此事
func (t T) M() {
	fmt.Println(t.S)
}

func Go1() {
	var a I = T{
		"草泥马",
	}

	a.M() //草泥马
}

/*
接口值

接口也是值。它们可以像其它值一样传递。

接口值可以用作函数的参数或返回值。

在内部，接口值可以看做包含值和具体类型的元组：
(value, type)

接口值保存了一个具体底层类型的具体值。

接口值调用方法时会执行其底层类型的同名方法。
*/

type Iphone interface {
	M()
}

type T1 struct {
	S string
}

func (t *T1) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func Go2() {
	var i Iphone

	i = &T1{"Hello"}
	describe(i) //(&{Hello}, *lession17.T1)
	i.M()

	i = F(math.Pi)
	describe(i) //(3.141592653589793, lession17.F)
	i.M()
}
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
