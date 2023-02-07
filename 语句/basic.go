package main

import "fmt"

// if语句
func iF(i int64) {
	if i > 100 {
		fmt.Print("大于100")
	} else if i > 50 && i <= 100 {
		fmt.Print("大于50小于等于100")
	} else {
		fmt.Println("小于等于50")
	}
}
func main() {
	iF(100)
}
