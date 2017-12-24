// 结构体文法通过直接列出字段来新分配一个结构体
// 使用 Name: 语法可以仅列出部分字段（字段名与顺序无关)
// 特殊的 & 前缀可以返回一个结构体的指针
package main

import "fmt"

type VerTex struct {
	X, Y int
}

func main() {
	v1 := VerTex{1,10}
	v2 := VerTex{Y:2}
	v3 := VerTex{}
	p := &VerTex{10, 2}

	fmt.Println(v1, v2, v3, p)
}