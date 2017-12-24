//信道可以带缓冲的
//缓冲长度可以作为第二个参数传给 make, 来初始化一个带缓冲的信道
package main

import "fmt"

func main(){
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}