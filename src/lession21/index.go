package lession21

import (
	"fmt"
	"image"
	"io"
	"strings"
)

/*
Reader

io 包指定了 io.Reader 接口，它表示从数据流的末尾进行读取。

Go 标准库包含了该接口的许多实现，包括文件、网络连接、压缩和加密等等。

io.Reader 接口有一个 Read 方法：

func (T) Read(b []byte) (n int, err error)

Read 用数据填充给定的字节切片并返回填充的字节数和错误值。
在遇到数据流的结尾时，它会返回一个 io.EOF 错误。

示例代码创建了一个 strings.Reader 并以每次 8 字节的速度读取它的输出。
*/

func Go() {
	r := strings.NewReader("Hello,Reader")

	b := make([]byte, 8)
	fmt.Println(r, b) //&{Hello,Reader 0 -1} [0 0 0 0 0 0 0 0]
	fmt.Println()
	fmt.Println()
	fmt.Println()
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

/*
图像
image 包定义了 Image 接口：

package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}

注意: Bounds 方法的返回值 Rectangle 实际上是一个 image.Rectangle，它在 image 包中声明。

（请参阅文档了解全部信息。）
https://go-zh.org/pkg/image/#Image

color.Color 和 color.Model 类型也是接口，但是通常因为直接使用预定义的实现 image.RGBA 和 image.RGBAModel 而被忽视了。这些接口和类型由 image/color 包定义。
*/
func Go4() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
