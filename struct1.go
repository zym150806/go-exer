// 一个结构体（struct） 就是一个字段的集合
package main

import "fmt"

type Vertex struct {
	X int
	Y float32
	Z int
}

func main() {
	fmt.Println(Vertex{1, 2, 3})
}



