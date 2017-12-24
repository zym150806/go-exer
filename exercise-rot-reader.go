//编写一个实现了 io.Reader 并从另一个 io.Reader 中读取数据的 rot13Reader ， 通过应用 rot13 代换密码对数据流进行修改。
//
//rot13Reader 类型已经提供。实现 Read 方法以满足 io.Reader 。
package main

import (
	"io"
	"strings"
	"os"
)

type rot13Reader struct {
	r io.Reader
}

func Rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'M':
		b = b + 13
	case 'M' < b && b <= 'Z':
		b = b - 13
	case 'a' <= b && b <= 'm':
		b = b + 13
	case 'm' < b && b <= 'z':
		b = b - 13
	}
	return b
}

func (reader rot13Reader) Read(b []byte) (n int, err error) {
	n, err = reader.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = Rot13(b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}