package main

import (
	"fmt"
)

type VerTex struct {
	X, Y float64
}

func (v *VerTex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *VerTex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := VerTex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &VerTex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
