// 类型推导
// 在声明一个变量而不指定其类型时，变量的类型由右值推导得出
// 右值声明了类型时，新变量与其类型相同
// 右值未声明类型时，新变量的类型取决于常量的精度
package main

import "fmt"

func main() {
	v := 1 + 0.5i

	fmt.Printf("v is of type %T\n", v);
}
