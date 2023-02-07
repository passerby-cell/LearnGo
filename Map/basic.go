package main

import "fmt"

func main() {
	m1 := map[string]string{
		"name":    "go",
		"version": "1.20",
	}
	//定义空Map
	m2 := make(map[string]int)
	var m3 map[string]int
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	if val, isExist := m1["nama"]; isExist {
		fmt.Println(val)
	} else {
		fmt.Println("not exist")
	}
	delete(m1, "name")
	fmt.Println(m1)
}
