package main

import "fmt"

type VerTex struct{
	Lat, Long float64
}

var m = map[string]VerTex{
	"abc": {12.44, -0.325},
	"aaa": {12.34, -0.22},
}

func main() {
	fmt.Println(m)
}