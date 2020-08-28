# Go 语言核心编程

### Go 学习方法

1. 高效而愉快的学习
2. 先建立整体矿建，然后细节
3. 在实际工作中要培养用到什么，能够快速学习的能力
4. 先 know how 再 know why
5. 软件编程是一门 ‘做中学’ 的学科，不是会了再做，而是做了再会

## Golang 的概述

简介：

Go 语言保证了既能到达**静态编程语言的安全和性能**，又能达到**动态语言维护的高效
率**

1. 从 c 语言中继承了很多理念，包括表达式语法，控制结构，基础数据类型，调用参数传
   值，指针等，也保留了和 c 一样的编译执行方式弱化的指针。

   ```go
   func testPtr(num *int){
       *num = 20
   }
   ```

2. 引入**包的概念**，用于组织程序结构，Go 语言的一个文件都要归属于一个包，不能单
   独存在

3. 垃圾回收机制，内存自动回收，不需要开发人员管理

4. **天然并发**

   - 从语言层面支持并发，实现简单
   - goroutine，轻量级线程，可实现大并发处理，高效利用多核
   - 基于 CPS 并发模型（Communicating Sequential Processes)实现

5. 吸收了管道通信机制，形成了 Go 语言特有的管道 channel，通过管道 Channel，可以
   实现不同的 goroutine 之间的相互通信

6. 函数返回多个值

7. 新的创新：切片、延时执行 defer 等

### 开发工具

vscode ，可以运行再 Mac OS X、windows 、linux 上 ，\*\*默认支持 Go 语言高亮，安
装 Go 插件，还智能提示。

安装步骤：

1. 安装 vscode

2. 安装 Go 的 SDK

![image-20200827073450998](D:\Technology
stack\GoGoGo\image-20200827073450998.png)

3. 配置 环境变量

| 环境变量 | 说明                             |
| -------- | -------------------------------- |
| GOROOT   | 指定 sdk 安装路径，C:\go         |
| Path     | 添加 SDK/bin 目录                |
| GOPATH   | 工作目录，将来 go 项目的工作路径 |

### 快速开发入门

写一个 hello.go 程序，可以输出 hello world

#### go 程序的目录结构

```shell
d:goproject
 -src
  ---go_code
    ----project01
    ----project02
       ------main
       和其他包
```

![image-20200827075826831](D:\Technology
stack\GoGoGo\image-20200827075826831.png)

#### hello.go

代码:

```go
//要求开发一个hello.go程序，可以输出'hello world'
package main

import "fmt"

func main(){
	fmt.Priny("Hello world")
}
```

说明：

- go 文件的后缀时.go
- `package main` 表示该 hello.go 文件所在的包时 main，在 go 中，每个文件都必须归
  属于一个包。
- `import "fmt"` 表示引入一个包，包名为 `fmt` ，引入该包后就可以使用`fmt`包里面
  的函数，比如：fmt.Println

- `func`是一个关键字，表示一个函数
- `main`时函数名，是一个主函数，程序的入口

通过 go build 命令对文件编译，生成.exe 文件，通过 go run 可以直接运行 hello.go

![image-20200827081013900](D:\Technology
stack\GoGoGo\image-20200827081013900.png)

> https://tour.golang.org/
>
> 暂时访问不了

### Golang 标准库 API 文档

中文网:

> studygolang.com/pkgdoc

## go 插件安装

查看哪些插件没有成功：

## 配置 go 的代理

> https://goproxy.io/zh/docs/getting-started.html



### 创建文件夹，打开项目

```powershell
PS D:\goplace> go mod init github.com/lubenweiniubi1/hellogo
go: creating new go.mod: module github.com/lubenweiniubi1/hellogo
PS D:\goplace> 

```

```go
module github.com/lubenweiniubi1/hellogo

go 1.15

```

问题：

```bash
$GOPATH/go.mod exists but should not
```

开启模块支持后，并不能与**GOPATH**共存GOPATH中移出即可

## Go tour

问题：

```shell
index.go:3:8: import "lession1" is a program, not an importable package
```

解决

本来所有go文件都在一个文件夹中，这里把go文件分成了coin和core两个文件夹。`coin和core文件目录中go文件设置了同样的packege main`，导致报错，**工具包的 package不能是main**

**你的包要放在src下面，并关闭mod模式**

关闭mod以后就可以在bin中运行tour.exe 了



