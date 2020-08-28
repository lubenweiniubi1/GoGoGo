package lession2

import (
	"fmt"
	"math"
)

/*
此代码用圆括号组合了导入，这是“分组”形式的导入语句。
当然你也可以编写多个导入语句，例如：
*/

// import "fmt"
// import "math"

/*
导出名
在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的。
例如，Pizza 就是个已导出名，Pi 也同样，它导出自 math 包。
pizza 和 pi 并未以大写字母开头，所以它们是未导出的。
在导入一个包时，你只能引用其中已导出的名字。任何“未导出”的名字在该包外均无法访问。
执行代码，观察错误输出。
然后将 math.pi 改名为 math.Pi 再试着执行一次。
*/

//lession2.lession2() //cannot refer to unexported name lession2.lession2
func lession2() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

//可以导出的
func Lession2() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
