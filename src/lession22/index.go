package lession22

import (
	"fmt"
	"time"
)

/*
Go 程

Go 程（goroutine）是由 Go 运行时管理的轻量级线程。

go f(x, y, z)

会启动一个新的 Go 程并执行
f(x, y, z)

f, x, y 和 z 的求值发生在当前的 Go 程中，而 f 的执行发生在新的 Go 程中。

Go 程在相同的地址空间中运行，
因此在访问共享的内存时必须进行同步。
sync 包提供了这种能力，不过在 Go 中并不经常用到，因为还有其它的办法（见下一页）。
*/
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func Go() {
	// say("world")
	// say("hello")
	//去掉go
	// world
	// world
	// world
	// world
	// world
	// hello
	// hello
	// hello
	// hello
	// hello

	go say("hello")
	// go say("world") 会两个都不执行
	say("world") //这里必须接一个不go的say
}

/*
信道

信道是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值。

ch <- v    // 将 v 发送至信道 ch。
v := <-ch  // 从 ch 接收值并赋予 v。
（“箭头”就是数据流的方向。）

和映射与切片一样，信道在使用前必须创建：

ch := make(chan int)
默认情况下，发送和接收操作在另一端准备好之前都会阻塞。

这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。

以下示例对切片中的数进行求和，将任务分配给两个 Go 程。
一旦两个 Go 程完成了它们的计算，它就能算出最终的结果。
*/
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println(sum)
	c <- sum //将和送入c
}

func Go1() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	go sum(s[:len(s)/2], c) //
	go sum(s[len(s)/2:], c) //这个总是先执行 ，可能跟栈有关

	x, y := <-c, <-c //从c中接收数据
	fmt.Println(x, y, x+y)
	fmt.Println(s)
}

/*
带缓冲的信道

信道可以是 带缓冲的。将缓冲长度作为第二个参数提供给 make 来初始化一个带缓冲的信道：
ch := make(chan int, 100)

仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。

修改示例填满缓冲区，然后看看会发生什么。
*/
func f2(i int, c chan int) {
	c <- i
}

func Go2() {
	// ch := make(chan int, 1) //fatal error: all goroutines are asleep - deadlock!
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	// go f2(1, ch)
	// go f2(2, ch) //多线程 缓冲区为0，或者不写 就可以塞数据了
	fmt.Println(<-ch) //1
	fmt.Println(<-ch) //2
}

/*
range 和 close

发送者可通过 close 关闭一个信道来表示没有需要发送的值了。
接收者可以通过为接收表达式分配第二个参数来测试信道是否被关闭：
若没有值可以接收且信道已被关闭，那么在执行完
v, ok := <-ch

此时 ok 会被设置为 false。

循环 for i := range c 会不断从信道接收值，直到它被关闭。

注意： 只有发送者才能关闭信道，而接收者不能。
向一个已经关闭的信道发送数据会引发程序恐慌（panic）。

还要注意： 信道与文件不同，通常情况下无需关闭它们。
只有在必须告诉接收者不再有需要发送的值时才有必要关闭，
例如终止一个 range 循环。
*/
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // 去掉会出现deadlock
}

func Go3() {
	c := make(chan int, 10)
	b := make(chan int, 10)

	b <- 1

	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

//go通道缓冲 使我们可以异步的读写通道
/*
默认情况下 通道是不带缓冲的
发送端发送数据
同时必须又接收端相应的接收数据
而带缓冲区的通道则允许发送端的数据发送和接收端的数据获取处于异步状态
就是说发送端发送的数据可以防范缓冲区里面 可以等待接收端取获取数据
而不是立刻需要接收端去获取数据
不过由于缓冲区的大小是有限的
所以还是必须有接收端来接收数据
否则缓冲区已满
数据发送端就 无法再发送数据了
*/

/*
select 语句

select 语句使一个 Go 程可以等待多个 通信chan 操作。

select 会阻塞到某个分支可以继续执行为止，
这时就会执行该分支。
当多个分支都准备好时会随机选择一个执行。
*/
func feb(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: //将值塞入c中
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func Go4() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ { //一开始会等待case 1，等他塞值进来
			fmt.Println(<-c) //然后从c中取出值10次，并打印，如果c中没有值，就等待 chan 塞进来
		}
		quit <- 0
	}()
	feb(c, quit)
}

/*
默认选择

当 select 中的其它分支都没有准备好时，default 分支就会执行。

为了在尝试发送或者接收时不发生阻塞，可使用 default 分支：

select {
case i := <-c:
    // 使用 i
default:
    // 从 c 中接收会阻塞时执行
}
*/
func Go5() {
	tick := time.Tick(100 * time.Millisecond)  //每 100ms 产出一坨东西
	boom := time.After(500 * time.Millisecond) //500ms 后产出一坨东西
	for {
		select {
		case <-tick:
			fmt.Println("tick.") //产出后会tick以下
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond) //默认的是50ms 打印   .
		}
	}

}
