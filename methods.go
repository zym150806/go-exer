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

func main(){
	v := VerTex{3, 4}
	fmt.Println(v.Abs())
}
