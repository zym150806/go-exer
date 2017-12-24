//for 循环的 range 形式可以遍历切片或映射
//使用for循环时，会返回两个值：当前元素的下标、对应下标的元素的副本
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}