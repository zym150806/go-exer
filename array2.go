//数组[n]T 表示拥有n个T类型的值的数组
//表达式：var a [10]int，会将变量a声明为拥有10个整数的数组
//go中数组大小是其类型的一部分，因此不能改变数组的大小
package main

import "fmt"

func main() {
	var caomingrui [2]string
	caomingrui[0] = "hello"
	caomingrui[1] = "xiaobaobao"
	fmt.Println(caomingrui[0], caomingrui[1])
	fmt.Println(caomingrui)

	permise := [2]string{"hello", "hehe"}
	fmt.Println(permise)
}