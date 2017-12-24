package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case bool:
		fmt.Printf("boolean value is %v", v)
	default:
		fmt.Printf("Whoops don't know")
	}
}

func main() {
	do(21)
	do("world")
	do(true)
}