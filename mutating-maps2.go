//通过双赋值检测某个键是否存在：v, ok := elem[key]
//若 key 在 elem 中，则ok为true，否则为false
//若 key 不在映射中，则 elem 为该映射元素类型的零值
//当从映射中读取不存在的值时，返回值为该映射元素类型的零值
package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 332
	fmt.Println("the value: ", m["Answer"])

	m["Answer"] = 45
	fmt.Println("the value: ", m["Answer"])

	delete(m, "Answer")
	fmt.Println("the value: ", m["Anser"])

	v, ok := m["Answer"]
	fmt.Println("the value: ", v, " Present? ", ok)
}