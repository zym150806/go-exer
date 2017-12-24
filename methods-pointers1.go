package main

import (
	"math"
	"fmt"
)

type VerTex struct {
	X, Y float64
}

func (v VerTex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *VerTex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := VerTex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}