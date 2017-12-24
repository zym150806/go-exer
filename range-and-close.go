//发送者可以通过close关闭一个信道来表示没有需要发送的值了；
//接收者可以通过为接收表达式分配第二个参数来测试是否被关闭：v, ok := <-ch
//注：
//1.只有发送者才能关闭信道；
//2.向一个已关闭的信道发送数据会引发panic
//3.信道与文件不同，通常情况下无需关闭
package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}