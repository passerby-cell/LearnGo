package main

import (
	"fmt"
	"strconv"
)

/*
总结：
	1、for、if后面的条件没有括号
	2、if条件里面也可以定义变量
	3、没有while
	4、switch没有break（自动break），也可以switch多个条件
*/
// if语句
func iF(i int64) {
	if i > 100 {
		fmt.Print("大于100\n")
	} else if i > 50 && i <= 100 {
		fmt.Print("大于50小于等于100\n")
	} else {
		fmt.Println("小于等于50\n")
	}
}

// switch语句
func sWitch(score int) {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score <= 80:
		g = "C"
	case score <= 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("不合要求\n"))
	}
	fmt.Printf("%s\n", g)
}

// if语句
func fOR(n int) string {
	binary := ""
	for ; n > 0; n = n / 2 {
		lsb := n % 2
		binary = strconv.Itoa(lsb) + binary
	}
	return binary
}

// 死循环
func forALL() {
	for {
		fmt.Println("abc")
	}

}
func main() {
	iF(100)
	sWitch(80)
	fmt.Println(fOR(5))
	//forALL()
}
