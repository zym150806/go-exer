package main

import "fmt"

type VerTex struct{
	Lat, Long float64
}

var m = map[string]VerTex{
	"ccc": {12.33, -77.11},
	"abcc": {11.22, -0.03},
}

func main() {
	fmt.Println(m)
}
