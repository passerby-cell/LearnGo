package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type returnFib func() int

func (r returnFib) Read(p []byte) (n int, err error) {
	next := r()
	if next > 1000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// Fibonacci数列
func fibonacci() returnFib {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
func main() {
	a := fibonacci()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(a())
	//}
	printFileContents(a)
}
