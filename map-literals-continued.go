package main

import "fmt"

type VerTex struct{
	Lat, Long float64
}

var m = map[string]VerTex{
	"abb": {12.34, -33.44},
	"kuka": {33.22, -12.22},
}

func main() {
	fmt.Println(m)
}