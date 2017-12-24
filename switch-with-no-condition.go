// 没有条件的switch和switch true一样
// 这种形式能将长串的if-then-else写的更清晰
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch{
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening!")
	}
}

