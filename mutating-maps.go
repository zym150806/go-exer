//通过双赋值检测某个键是否存在 ：elem, ok = elem[key];
//若 key 在 elem中，则 ok 为 true，否则 ok 为 false。
//若 key 不在映射中，则 elme 为该映射元素类型的零值
//当从映射中读取某个不存在的键时，返回的值为映射元素类型的零值；
package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Anwser"] = 42
	fmt.Println("the value: ", m["Anwser"])

	m["Anwser"] = 45
	fmt.Println("the value: ", m["Anwser"])

	delete(m, "Anwser")
	fmt.Println("the value: ", m["Anwser"])

	v, ok := m["Anwser"]
	fmt.Println("The value: ", v , ", Present: ", ok)
}