package main

import "fmt"

func main() {
	str := "ABCDCBA"
	ans := RecursivePalindrome(str, 0, len(str)-1)
	if ans {
		fmt.Printf("%s is palindrome\n", str)
	} else {
		fmt.Printf("%s is not palindrome\n", str)
	}
}

func RecursivePalindrome(arr string, start, end int) bool {
	if start >= end {
		return true
	}
	if arr[start] != arr[end] {
		return false
	}
	return RecursivePalindrome(arr, start+1, end-1)
}
