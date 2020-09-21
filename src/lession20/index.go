package lession20

import (
	"fmt"
	"math"
	"time"
)

/*
Stringer

fmt 包中定义的 Stringer 是最普遍的接口之一。

type Stringer interface {
    String() string
}

Stringer 是一个可以用字符串描述自己的类型。fmt 包（还有很多包）都通过此接口来打印值。
*/

func Go() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod  Beeblebrox", 9001}
	fmt.Println(a, z) //Arthur Dent (42 years) Zaphod  Beeblebrox (9001 years)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

/*
练习：Stringer

通过让 IPAddr 类型实现 fmt.Stringer 来打印点号分隔的地址。

例如，IPAddr{1, 2, 3, 4} 应当打印为 "1.2.3.4"。
*/

type IPAddr [4]byte

// TODO: 给 IPAddr 添加一个 "String() string" 方法

func (addr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v\n", addr[0], addr[1], addr[2], addr[3])
}

func Practice() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

/*
错误

Go 程序使用 error 值来表示错误状态。

与 fmt.Stringer 类似，error 类型是一个内建接口：
type error interface {
    Error() string
}

（与 fmt.Stringer 类似，fmt 包在打印值时也会满足 error。）

通常函数会返回一个 error 值，调用的它的代码应当判断这个错误是否等于 nil 来进行错误处理。
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
error 为 nil 时表示成功；非 nil 的 error 表示失败。
*/
func Go2() {
	if err := run(); err != nil {
		fmt.Println(err) //at 2020-09-01 11:00:31.3001423 +0800 CST m=+0.003061601 ,it didn't work
	}
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

type MyError struct {
	When time.Time
	What string
}

func (m *MyError) Error() string {
	return fmt.Sprintf("at %v ,%s", m.When, m.What)
}

/*
练习：错误

从之前的练习中复制 Sqrt 函数，修改它使其返回 error 值。

Sqrt 接受到一个负数时，应当返回一个非 nil 的错误值。复数同样也不被支持。

创建一个新的类型
type ErrNegativeSqrt float64

并为其实现
func (e ErrNegativeSqrt) Error() string
方法使其拥有 error 值，
通过 ErrNegativeSqrt(-2).Error() 调用该方法应返回 "cannot Sqrt negative number: -2"。

注意: 在 Error 方法内调用 fmt.Sprint(e) 会让程序陷入死循环。

可以通过先转换 e 来避免这个问题：fmt.Sprint(float64(e))。这是为什么呢？
修改 Sqrt 函数，使其接受一个负数时，返回 ErrNegativeSqrt 值。
*/
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	return math.Sqrt(x), nil
}

func Practice2() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
