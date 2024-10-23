package main

import "fmt"

func main() {
	arr := []int{5, 6, 7, 88}
	fmt.Printf("Largest Element: %v", LargestElement(arr, 0, arr[0]))
}

func LargestElement(arr []int, index, op int) int {
	if index > len(arr)-1 {
		return op
	}
	if op < arr[index] {
		op = arr[index]
	}
	return LargestElement(arr, index+1, op)
}
