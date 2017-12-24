//数组 [n]T 表示拥有n个T类型的值的数组
//表达式： var a [10]int, 会将变量a声明为有10个整数的数组
//数组大小是其类型的一部分，因此数组不能改变大小
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(primes)
}
