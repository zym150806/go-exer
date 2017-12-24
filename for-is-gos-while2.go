package main

import "fmt"

func main() {
	sum := 1
	for sum < 10000 {
		fmt.Println(sum)
		sum += sum
	}

	fmt.Println(sum)
}
