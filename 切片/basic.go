package main

import "fmt"

func change(arr []int) {
	arr[0] = 100
}
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//fmt.Println("arr[2:6]=", arr[2:6])
	//fmt.Println("arr[:6]=", arr[:6])
	//fmt.Println("arr[2:]=", arr[2:])
	//fmt.Println("arr[:]=", arr[:])
	//change(arr[:])
	//fmt.Println(arr)

	s1 := arr[2:5]
	fmt.Println(s1)

	s2 := s1[2:5]
	fmt.Println(s2)

	s3 := append(s2, 10)
	fmt.Println(s3)
	s4 := append(s3, 11)
	fmt.Println(s4)
	s5 := append(s4, 12)
	fmt.Println(s5)
	fmt.Println(arr)

}
