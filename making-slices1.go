//make 用于创建切片和动态数组
//make 函数会分配一个元素为零值的数组并返回一个引用了它的切片
//make 函数可以通过传入第三个参数的方式来指定容量
//b := make([]int, 0, 5)
package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}