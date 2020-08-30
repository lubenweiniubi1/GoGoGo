package lession12

import "fmt"

/*
映射 (`map`)

映射将键映射到值。

映射的零值为 nil 。nil 映射既没有键，也不能添加键。

make 函数会返回给定类型的映射，并将其初始化备用。
*/

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func Go() {
	fmt.Println(m) //map[]
	m = make(map[string]Vertex)
	fmt.Println(m) //map[]

	m["BellLabs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m)             //map[BellLabs:{40.68433 -74.39967}]
	fmt.Println(m["BellLabs"]) //{40.68433 -74.39967}

}

/*
映射的文法

映射的文法与结构体相似，不过必须有键名。
*/
func Go1() {
	var m = map[string]Vertex{
		"BellLabs": Vertex{
			40.68452, -74111,
		},
		"Google": Vertex{
			37, -122,
		},
	}
	fmt.Println(m) //map[BellLabs:{40.68452 -74111} Google:{37 -122}]

}

//若顶级类型只是一个类型名，你可以在文法的元素中省略它。
func Go2() {
	var m = map[string]Vertex{
		"Bell Labs": {112, 112},
	}
	fmt.Println(m) //map[BellLabs:{40.68452 -74111} Google:{37 -122}]

}

/*
在映射 m 中插入或修改元素：
m[key] = elem

获取元素：
elem = m[key]

删除元素：
delete(m, key)

通过双赋值检测某个键是否存在：
elem, ok = m[key]

若 key 在 m 中，ok 为 true ；否则，ok 为 false。
若 key 不在映射中，那么 elem 是该映射元素类型的零值。

同样的，当从映射中读取某个不存在的键时，结果是映射的元素类型的零值。
注 ：若 elem 或 ok 还未声明，你可以使用短变量声明：
elem, ok := m[key]

*/
func Go3() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"]) //The value: 42

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"]) //The value: 48

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok) //The value: 48 Present? true

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"]) //The value: 0

	v, ok = m["Answer"]
	fmt.Println("The value:", v, "Present?", ok) //The value: 0 Present? false

}
