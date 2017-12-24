//切片拥有长度和容量
//切片的长度就是它所包含的元素个数
//切片的容量从它第一个元素数，到底层数组元素末位的个数
//切片s的长度和容量可以分别通过 len(s) 和 cap(s) 来获取
//可以通过重新切片的方式来扩展一个切片，从而给它提供足够的容量
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 9, 11, 13}
	printSlice(s)

	// slice the slice to give it zero length
	s = s[:0]
	printSlice(s)

	// extends its length
	s = s[:4]
	printSlice(s)

	// drop its first two value
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}