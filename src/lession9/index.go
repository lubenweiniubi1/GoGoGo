package lession9

import "fmt"

//一个结构体（struct）就是一组字段（field）。

type Vertex struct {
	X int
	Y int
}

func Go() {
	fmt.Println(Vertex{2, 11}) //{2 11}
}

//结构体字段使用点号来访问。
func Go1() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)   //{4 2}
	fmt.Println(v.X) //4
}

/*
结构体指针

结构体字段可以通过结构体指针来访问。

如果我们有一个指向结构体的指针 p，那么可以通过 (*p).X 来访问其字段 X。
不过这么写太啰嗦了，所以语言也允许我们使用隐式间接引用，直接写 p.X 就可以。
*/

type V struct {
	X int
	Y int
}

func Go2() {
	v := V{1, 2}
	p := &v
	p.X = 777
	fmt.Println(v) //{777 2}
}

/*
结构体文法

结构体文法通过直接列出字段的值来新分配一个结构体。

使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）

特殊的前缀 & 返回一个指向结构体的指针。
*/
type Ver struct {
	X, Y int
}

var (
	v1 = Ver{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Ver{X: 1}  // Y:0 被隐式地赋予
	v3 = Ver{}      // X:0 Y:0
	p  = &Ver{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)

func Go3() {
	fmt.Println(v1, v2, v3, p) //{1 2} {1 0} {0 0} &{1 2}
}
