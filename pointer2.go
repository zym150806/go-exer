// go 拥有指针类型，指针保存了值的内存地址
// 类型 *T 是指向 T 类型的指针，指针的零值是nil
// & 操作符会生成一个指向其操作数的指针
// * 操作符表示指针指向底层的值，即经常说的"简介引用"或"重定向"
// go 语言没有指针运算
package main

import "fmt"

func main() {
	i, j := 10, 20
	var tmp *int

	fmt.Println(tmp)

	p := &i // point to i
	fmt.Println(p)
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p += 2
	fmt.Println(j)


}