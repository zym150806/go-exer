package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	sort := 1
	val1 := 0
	val2 := 0
	temp := 0
	return func() int {
		if sort == 1 {
			val1 = 0
		}

		if sort == 2 {
			val2 = 1
			temp = 1
		}

		if sort > 2 {
			temp = val1 + val2
			val1 = val2
			val2 = temp
		}
		sort++

		return temp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
