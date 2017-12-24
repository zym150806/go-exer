//数组[n]T 表示拥有n个T类型的值的数组
//表达式：var a [10]int ,会将变量a声明为有10个整数的数组
//go中数组大小是其类型的一部分，因此不能改变数组的大小
package main

import "fmt"

func main() {
	var hello [2]string

	hello[0] = "Hello"
	hello[1] = "World"
	fmt.Println(hello[0], hello[1])
	fmt.Println(hello)

	primes := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(primes)
}