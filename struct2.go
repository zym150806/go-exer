// 一个结构体（struct）就是一个字段的集合
package main

import "fmt"

type VerTex struct {
	X int
	Y int
	Z int
}

func main() {
	fmt.Println(VerTex{3.3, 2.2, 1.1})
}