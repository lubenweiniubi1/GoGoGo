package lession10

import "fmt"

/*
类型 [n]T 表示拥有 n 个 T 类型的值的数组。

表达式

var a [10]int
会将变量 a 声明为拥有 10 个整数的数组。

数组的长度是其类型的一部分，因此数组不能改变大小。
这看起来是个限制，不过没关系，Go 提供了更加便利的方式来使用数组。
*/
func Go() {
	var a [2]string
	a[0] = "hello"
	a[1] = " world"
	fmt.Println(a[0], a[1]) //hello  world
	fmt.Println(a)          //[hello  world]

	primes := [6]int{2, 3, 5, 7, 11, 13} //初始化
	fmt.Println(primes)                  //[2 3 5 7 11 13]
}

/*
切片

每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角。
在实践中，切片比数组更常用。

类型 []T 表示一个元素类型为 T 的切片。

切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔：

a[low : high]
它会选择一个半开区间，包括第一个元素，但排除最后一个元素。

以下表达式创建了一个切片，它包含 a 中下标从 1 到 3 的元素：

a[1:4]
*/
func Go1() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[0:4]
	var s1 []int = primes[3:4]
	fmt.Println(s, s1) //[2 3 5 7] [7]
}

/*
切片就像数组的引用
切片并不存储任何数据，它只是描述了底层数组中的一段。

更改切片的元素会修改其底层数组中对应的元素。

与它共享底层数组的切片都会观测到这些修改。
*/
func Go2() {
	names := [4]string{
		"张三",
		"李四",
		"王五",
		"赵六",
	}
	fmt.Println(names) //[张三 李四 王五 赵六]

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b) //[张三 李四] [李四 王五]

	b[0] = "吴一凡"
	fmt.Println(a, b) //[张三 吴一凡] [吴一凡 王五]
}

/*
切片文法

切片文法类似于没有长度的数组文法。

这是一个数组文法：

[3]bool{true, true, false}

下面这样则会创建一个和上面相同的数组，然后构建一个引用了它的切片：

[]bool{true, true, false}
*/
func Go3() {
	q := []int{2, 3, 4, 5}
	fmt.Println(q) //[2 3 4 5 6 7 8 9]

	r := []bool{true, false, true, true, false, true} //[true false true true false true]
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{4, true},
		{5, false},
	}
	fmt.Println(s)      //[{2 true} {3 false} {4 true} {5 false}]
	fmt.Println(q[0:1]) // [2]
}

/*
切片的默认行为

在进行切片时，你可以利用它的默认行为来忽略上下界。

切片 下界的默认值为 0，上界则是该切片的长度。

对于数组

var a [10]int
来说，以下切片是等价的：

a[0:10]
a[:10]
a[0:]
a[:]
*/
func Go4() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)

	s = s[1:4]
	fmt.Println(s) //[3 5 7]
	printSlice(s)  //len=3 cap=5 [3 5 7]

	s = s[:2]
	fmt.Println(s) //[3 5]

	s = s[1:]
	fmt.Println(s) //[5]

	s = s[:]
	fmt.Println(s) //[5]

	s = s[:4] //拓宽回来

	fmt.Printf("在这里切片儿: %v", s) //  [5 7 11 13]
}

/*
切片的长度与容量

切片拥有 长度 和 容量。

切片的长度就是它所包含的元素个数。

切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。 跟一般容量概念一样

切片 s 的长度和容量可通过表达式 len(s) 和 cap(s) 来获取。

只要具有足够的容量，你就可以通过重新切片来扩展一个切片。
请试着修改示例程序中的某个切片操作，使其长度超过容量（即将它扩展到超出其容量范围），看看会发生什么。
*/

func Go5() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) //len=6 cap=6 [2 3 5 7 11 13]

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s) //len=0 cap=6 []

	// 拓展其长度
	s = s[:4]
	printSlice(s) //len=4 cap=6 [2 3 5 7]

	// 舍弃前两个值
	s = s[2:]
	printSlice(s) //len=2 cap=4 [5 7]

	if false {
		// 长度超出容量
		s = s[0:5] //报错
		printSlice(s)
	}

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

/*
nil 切片

切片的零值是 nil。

nil 切片的长度和容量为 0 且没有底层数组。
*/

func Go6() {
	var s []int
	fmt.Println(s, len(s), cap(s)) //[] 0 0
	if s == nil {
		fmt.Println("nil")
	}
}

/*
用 make 创建切片

切片可以用内建函数 make 来创建，这也是你创建动态数组的方式。

make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：

a := make([]int, 5)  // len(a)=5

要指定它的容量，需向 make 传入第三个参数：

b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
*/
func Go7() {
	a := make([]int, 5)
	printSlice(a) //len=5 cap=5 [0 0 0 0 0]

	b := make([]int, 0, 5)
	printSlice(b) //len=0 cap=5 []

	c := b[:2]
	printSlice(c) //len=2 cap=5 [0 0]

	d := c[2:5]
	printSlice(d) //len=3 cap=3 [0 0 0]

}

/*
切片的切片

切片可包含任何类型，甚至包括其它的切片。

下面代码 如果是
[] string {
	 []string{"_", "_", "_"},
     []string{"_", "_", "_"},
	 []string{"_", "_", "_"},
}
 打括号里面应该是 string类型的数据，现在却是 切片儿，所以报错，因为大括号里面是string的切片儿
 所以 数组类型因该是切片

 即   [10]-- []string ---{
	[]string{xxxxxx}
 }
*/
func Go8() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "x"
	board[0][2] = "x"
	board[2][2] = "o"
	board[1][2] = "x"
	board[1][0] = "o"

	fmt.Println(board)
}

/*
向切片追加元素
为切片追加新的元素是种常用的操作，为此 Go 提供了内建的 append 函数。内建函数
的文档
https://go-zh.org/pkg/builtin/#append
对此函数有详细的介绍。

func append(s []T, vs ...T) []T
append 的第一个参数 s 是一个元素类型为 T 的切片，其余类型为 T 的值将会追加到该切片的末尾。

append 的结果是一个包含原切片所有元素加上新添加元素的切片。

当 s 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。返回的切片会指向这个新分配的数组。

（要了解关于切片的更多内容，请阅读文章 Go 切片：用法和本质。
https://blog.go-zh.org/go-slices-usage-and-internals
）
*/

func Go9() {
	var s []int
	printSlice(s) //len=0 cap=0 []

	//添加一个空切片
	s = append(s, 0)
	printSlice(s) //len=1 cap=1 [0]

	//这个切片会按需增长
	s = append(s, 1)
	printSlice(s) //len=2 cap=2 [0 1]

	//可以一次性添加多个元素
	s = append(s, 2, 3, 4, 5)
	printSlice(s) //len=6 cap=6 [0 1 2 3 4 5]

}
