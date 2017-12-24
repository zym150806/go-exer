// go 语言的基本数据类型
// bool
// string
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64
//
// byte // uint8的别名
// rune // int32的别名，表示一个 Unicode 码点
// float32 float64
// complex64 complex128
//
// 同导入语句一样，变量声明也可以'分组'成一个语法块
// int uint uintptr 在32位系统上通常为32位宽，在64位系统上则为64位
// 当需要一个整数值时应使用 int 类型，除非有特殊理由使用固定大小或者无符号的整数类型
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	Tobe   bool = true
	MaxInt int64 = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", Tobe, Tobe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
