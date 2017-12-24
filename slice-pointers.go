//切片不存储数据，它只描述了底层数组的一段；
//更改切片的元素会修改其底层数组对应的元素
//与当前切片共享底层数组的切片都会观测到修改
package main

import "fmt"

func main() {
	names := [4]string{
		"lee",
		"kevin",
		"john",
		"paul",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "aaa"
	fmt.Println(a, b)
	fmt.Println(names)
}