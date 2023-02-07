package main

import "fmt"

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("不支持的运算符" + op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}
func main() {
	fmt.Println(eval(2, 4, "+"))
	fmt.Println(div(21, 4))
	q, r := div(21, 4)
	//下划线表示第二个值不接收
	//q, _ := div(21, 4)
	fmt.Println(q, r)
}
