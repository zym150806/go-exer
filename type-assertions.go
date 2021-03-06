//类型断言提供了访问接口底层具体值的方法
//t := i.(T)
//该语句断言接口值 i 保存了具体类型T，并将其底层类型为T的值赋予变量t
//若 i 并未保存 T 类型的值，该语句会触发一个恐慌
//
//类型断言返回两个值：其底层值、报告断言是否成功的布尔值

//注： t, ok = i.(T), 当 i 保存了一个T，此时ok 为 true， 否则为false。
//    无论是true 或者 false，都不会产生恐慌
package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)
	fmt.Println(f)
}