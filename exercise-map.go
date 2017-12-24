//实现 WordCount 。它应当返回一个映射，其中包含每个字符串 s 中“单词”的个数。
// 函数 wc.Test 会对此函数执行一系列测试用例，并输出成功还是失败。
//你会发现 strings.Fields 很有帮助。
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	sArr := strings.Fields(s)
	for _, val := range sArr {
		result[val] += 1
	}

	return result
}

func main() {
	wc.Test(WordCount)
}