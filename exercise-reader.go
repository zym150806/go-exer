//实现一个 Reader 类型，它产生一个 ASCII 字符 'A' 的无限流。
package main

import "github.com/Go-zh/tour/reader"

type MyReader struct {

}

// TODO: Add a Read([]byte)(int, error) method to MyReader
func (r MyReader) Read(b []byte) (n int, err error) {
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}