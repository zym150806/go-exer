// go 拥有指针类型，指针保存了值的内存地址
// 类型 *T是指向T类型的指针，指针的零值是nil
// & 操作符会生成一个指向其操作数的指针
// * 操作符表示指针指向的底层值，及经常说的"间接引用"或者"重定向"
// go 语言没有指针运算
package main

import "fmt"

func main() {
	i, j := 42, 2071

	p := &i // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21 // set i through the pointer
	fmt.Println(i) // set the new value of i

	p = &j // point to j
	*p = *p / 37 // divide j through the pointer
	fmt.Println(j) // see the new value of j
}