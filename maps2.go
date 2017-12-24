// 映射将键映射到值；
// 映射的零值是nil。nil映射既没有键，也不能添加键。
// make 函数会返回给定类型的映射，并将其初始化备用
package main

import "fmt"

type VerTex struct {
	Lat, Long float64
}

var m map[string]VerTex

func main() {
	m = make(map[string]VerTex)
	m["Bell Labs"] = VerTex{
		40.123, -30.255,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m)
}