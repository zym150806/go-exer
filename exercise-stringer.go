//通过让 IPAddr 类型实现 fmt.Stringer 来打印点号分隔的地址。
//
//例如，IPAddr{1, 2, 3, 4} 应当打印为 "1.2.3.4" 。
package main

import (
	"fmt"
	"bytes"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var buffer bytes.Buffer
	for _, piece := range ip {
		buffer.WriteString(fmt.Sprintf("%v", piece))
		buffer.WriteString(".")
	}

	return strings.TrimRight(buffer.String(), ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}