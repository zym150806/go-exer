package main

import "fmt"

type VerTex struct {
	X int
	Y int
}

func main() {
	v := VerTex{1, 3}
	fmt.Println(v.X)
	v.X = 4
	fmt.Println(v.X)
}