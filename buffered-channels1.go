// 信道是可以带缓冲的
//缓冲长度可以作为第二个参数传给 make， 来初始化一个带缓冲的信道
package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	for i := 0; i < 5; i++ {
		ch <- i
	}

	fmt.Println(<-ch)
}
