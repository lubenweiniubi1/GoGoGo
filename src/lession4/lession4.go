package lession4

/*
变量

var 语句用于声明一个 变量列表，跟函数的参数列表一样，类型在最后。

就像在这个例子中看到的一样，var 语句可以出现在包或函数级别。

*/
import "fmt"

var c, python, java bool

func Lession4() {
	var i int
	fmt.Println(i, c, python, java) // 0 false false false ,这里bool类型的默认值为false，int类型默认值为0
}

/*
变量的初始化

变量声明可以包含初始值，每个变量对应一个。

如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。
*/
func Lession4_1() {
	var i, j int = 1, 2
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java) //1 2 true false no!
}

/*
短变量声明

在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明。

在函数外，每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用。
*/
func Lession4_2() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "!no"
	fmt.Println(i, j, k, c, python, java)

}
