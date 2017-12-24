// 结构体字段可以通过结构体指针来访问
package main

import "fmt"

type VerTex struct {
	X int
	Y int
}

func main() {
	v := VerTex{1, 4}
	p := &v
	fmt.Println(p.X)
	fmt.Println((*p).Y)
}