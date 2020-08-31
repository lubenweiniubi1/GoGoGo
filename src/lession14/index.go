package lession14

import "fmt"

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

func Lession8() {
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
