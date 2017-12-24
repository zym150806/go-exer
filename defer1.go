// defer 语句会将函数推迟到外层函数返回后执行
// 推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用
package main

import (
	"fmt"
)

func main() {
	sum := 1
	defer fmt.Print(sum+1)
	fmt.Println("sum value is now ", sum+1)
}
