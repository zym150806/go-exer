//当 select 中的其他分支都没准备好时， default 分支就会运行
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.Tick(500 * time.Millisecond)

	for {
		select {
			case <-tick:
				fmt.Println("tick.")
			case <-boom:
				fmt.Println("boom.")
				return
			default:
				fmt.Println("    .")
				time.Sleep(50 * time.Millisecond)
		}
	}
}