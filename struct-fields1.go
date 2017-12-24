package main

import "fmt"

type VerTex struct {
	X int
	Y int
}

func main() {
	v := VerTex{1, 5}
	fmt.Println(v.Y)
	v.X = 2
	fmt.Println(v.X)
}