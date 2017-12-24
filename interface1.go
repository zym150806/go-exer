//接口是由一组方法签名定义的集合
//接口类型的值可以保存任何实现了这些方法的值
package main

import (
	"math"
	"fmt"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := VerTex{3, 4}

	a =  f // a MyFloat 实现了 Abser
	fmt.Println(a.Abs())

	a = &v // a *VerTex 实现了 Abser

	// 下面一行，v是一个VerTex(而不是*VerTex)
	// 所以没有实现 Abser
	// a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

type VerTex struct {
	X, Y float64
}

func (v *VerTex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
