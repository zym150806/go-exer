//在进行切片时，可以用切片默认行为来忽略上下界
//切片默认下界为0， 默认上界为切片的长度
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 9, 11, 13}
	fmt.Println(s)

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}