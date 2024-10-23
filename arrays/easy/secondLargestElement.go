package main

import "fmt"

func main() {
	arr := []int{5, 6, 7, 88}
	fmt.Printf("Largest Element: %v", SecondLargestElement(arr, 0, arr[0], arr[0]))
}

func SecondLargestElement(arr []int, index, first, second int) int {
	if index > len(arr)-1 {
		return second
	}
	if arr[index] > first {
		first, second = arr[index], first
	} else if first < arr[index] && arr[index] > second {
		second = arr[index]
	}
	return SecondLargestElement(arr, index+1, first, second)
}
