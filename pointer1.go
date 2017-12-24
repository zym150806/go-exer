// go 拥有指针类型，指针保存了值的内存地址
// 类型 *T 是指向 T 类型的指针，指针的零值是nil
// & 操作符会生成一个指向其操作数的指针
// * 操作符表示指针指向的底层值，即经常说的"间接引用"或者"重定向"
// go 语言没有指针运算
package main

import "fmt"

func main() {
	i, j := 21, 2071
	var pt *int


	fmt.Println(pt) // empty pointer

	p := &i // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21 // set new value through pointer
	fmt.Println(i) // see the new value of i

	p = &j // point to j
	*p = *p / 37 // divide j through the pointer
	fmt.Println(j) // see the new value of j

}