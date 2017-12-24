// 结构体字段可以通过结构体指针来访问
package main

import "fmt"

type VerTex struct {
	X int
	Y int
}

func main() {
	v := VerTex{1,10}
	p := &v

	p.X = 1e2
	fmt.Println(v);
}