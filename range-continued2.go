//for循环的range形式，可以将不需要的索引或者值使用"_"来忽略掉
//若只需要索引，只要将value部分忽略掉即可
package main

import "fmt"

func main() {
	pow := make([]int, 10)
	fmt.Println(pow)

	for i := range pow {
		pow[i] = 1 << uint(i) // 2**i
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}