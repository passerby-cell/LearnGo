package main

import (
	"fmt"
	"math"
)

const fileName = "b.txt"
const (
	a = 1
	b = 2
)

// 常量定义
func consts() {
	const fileName = "a.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(fileName, a, b, c)
}

// 枚举类型
func enums() {
	const (
		//使用iota自增值
		cpp = iota
		java
		python
		javascript
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
func main() {
	consts()
	enums()
}
