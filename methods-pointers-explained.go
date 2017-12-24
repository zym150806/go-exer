package main

import (
	"math"
	"fmt"
)

type VerTex struct {
	X, Y float64
}

func Abs(v VerTex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *VerTex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := VerTex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
