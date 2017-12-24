//信道是带有类型的管道，可以通过信道操作符 <- 来发送或者接收值
//信道使用前必须创建 ch := make(chan int)
//发送和接收操作在另一端准备好之前都会阻塞
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从c中接收

	fmt.Println(x, y, x+y)
}