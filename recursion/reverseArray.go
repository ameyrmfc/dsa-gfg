package main

import "fmt"

func main() {
	arr := []int{10, 20, 30, 40, 50, 60, 70}
	fmt.Printf("Input arr : %v\n", arr)
	//fmt.Printf("Output arr : %v", ReverseArray(arr))
	RecursiveReverseArray(arr, 0, len(arr)-1)
	fmt.Printf("Recursive Output arr : %v", arr)
}

func ReverseArray(arr []int) []int {
	left := 0
	right := len(arr) - 1
	for left < right {

		temp := arr[left]
		arr[left] = arr[right]
		arr[right] = temp

		left++
		right--
	}
	return arr
}

func RecursiveReverseArray(arr []int, start, end int) {
	if start >= end {
		return
	}
	temp := arr[start]
	arr[start] = arr[end]
	arr[end] = temp
	RecursiveReverseArray(arr, start+1, end-1)
	return
}
