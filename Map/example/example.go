package main

import "fmt"

/*
给定一个字符串，请你找出其中不含有重复字符的最长子串的长度。
abcabcbb -----> abc
*/
func lengthOfNoRepeatingSubString(s string) int {
	lastAppear := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, v := range []rune(s) {
		lastI, ok := lastAppear[v]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastAppear[v] = i
	}
	return maxLength
}
func main() {
	str := "我是你爸爸acrfvbgtyhn！"
	fmt.Println(lengthOfNoRepeatingSubString(str))
}
