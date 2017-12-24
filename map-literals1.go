package main

import "fmt"

type VerTex struct{
	Lat, Long float64
}

var m = map[string]VerTex{
	"haha": VerTex{
		12.3, 22.2,
	},
	"hello": VerTex{
		12.2, 11.1,
	},
}

func main() {
	fmt.Println(m)
}
