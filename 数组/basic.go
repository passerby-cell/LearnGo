package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3}
	//四行五列
	var grid [4][5]bool
	fmt.Println(arr1, arr2, arr3, grid)

	for _, v := range arr3 {
		fmt.Println(v)
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}
}
