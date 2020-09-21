package lession18

import "fmt"

/*
底层值为 nil 的接口值

即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。

在一些语言中，这会触发一个空指针异常，但在 Go 中通常会写一些方法来优雅地处理它
（如本例中的 M 方法）。

注意: 保存了 nil 具体值的接口其自身并不为 nil。
*/
func Go() {
	var i I
	var t *T

	i = t
	describe(i) //(<nil>, *lession18.T)
	i.M()       //<nil>
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil { //没有这个会报错
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
nil 接口值

nil 接口值既不保存值也不保存具体类型。

为 nil 接口调用方法会产生运行时错误，
因为接口的元组内并未包含能够指明该调用哪个 具体 方法的类型。
*/
func Go1() {
	var i I
	describe(i) //(<nil>, <nil>)
	// i.M() 报错
}

/*
空接口

指定了零个方法的接口值被称为 空接口：
interface{}

空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）

空接口被用来处理未知类型的值。例如，fmt.Print 可接受类型为 interface{} 的任意数量的参数。
*/
func Go2() {
	var i interface{}
	describeI(i) //(<nil>, <nil>)

	i = 42
	describeI(i) //(42, int)

	i = "hello world"
	describeI(i) //(hello world, string)

}
func describeI(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
