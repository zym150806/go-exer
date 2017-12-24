package main

import "fmt"

type VerTex struct {
	X int
	Y int
}

func main() {
	v := VerTex{4, 2}
	fmt.Println(v.X)
	fmt.Println(v.Y)
}
