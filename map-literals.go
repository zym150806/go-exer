package main

import "fmt"

type VerTex struct{
	Lat, Long float64
}

var m = map[string]VerTex{
	"abc": VerTex{
		123.44, -222.22,
	},
	"def": VerTex{
		111.11, -333.33,
	},
}

func main() {
	fmt.Println(m)
}
