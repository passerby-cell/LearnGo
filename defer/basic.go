package main

import (
	"bufio"
	"fmt"
	"learngo/defer/fib"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}
func writeFile(file string) {
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	defer writer.Flush()
	fib := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, fib())
	}
}

func main() {
	//tryDefer()
	writeFile("test.txt")
}
