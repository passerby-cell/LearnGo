package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 定义包变量 不可以使用“:”
var globalA = 1

var (
	aa = 1
	bb = "23"
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q \n", a, s)
}

func variableInitValue() {
	var a, b int = 1, 3
	var s string = "23"
	fmt.Printf("%d %d %q \n", a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, d = 1, "23", true, 2.30
	fmt.Println(a, b, c, d)
}

func variableShorter() {
	//简写var：使用“:”
	a, b, c, d := 1, "23", true, 2.30
	//再次赋值就不用使用“:”
	a = 2
	fmt.Println(a, b, c, d)
}

// 内建类型 复数 complex
func euler() {
	fmt.Println(
		cmplx.Exp(1i*math.Pi)+1,
		cmplx.Pow(math.E, 1i*math.Pi)+1)
}

// 强制类型转换
func triangle() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	println(c)
}
func main() {
	fmt.Println("hello world!")
	fmt.Println(globalA, aa, bb)
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()
	euler()
	triangle()
}
