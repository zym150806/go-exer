// 没被明确初始值的变量声明，会被赋予他们的零值
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
