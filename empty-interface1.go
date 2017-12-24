//指定了零个方法的接口值被称为空接口
//空接口可以保存任何类型的值（因为每种类型都至少实现了零个方法）
//空接口被用来处理未知类型的值
package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)", i,i)
}