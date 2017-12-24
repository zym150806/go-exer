package main

import "fmt"

type VerTex struct{
	Lat, Long float64
}

var m = map[string]VerTex{
	"abc": VerTex{
		123, 23.22,
	},
	"kuka": VerTex{
		12.33, 33.44,
	},
}

func main() {
	fmt.Println(m)
}
