package main

import "math"

// 强制类型转换
func triangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	println(c)
	return c
}
func main() {
	triangle(3, 4)
}
