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
	f, err := os.OpenFile(file, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("error creating file: ", err)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("error closing file:", err)
		}
	}(f)
	writer := bufio.NewWriter(f)
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			fmt.Println("error flushing file:", err)
		}
	}(writer)
	fib := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, fib())
	}
}

func main() {
	//tryDefer()
	writeFile("test.txt")
}
