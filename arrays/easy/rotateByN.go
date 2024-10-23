package main

import "fmt"

func Reverse(arr []int, start, end int) {
	for start <= end {
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start++
		end--
	}
}

func main() {
	k := 2
	arr := []int{5, 4, 3, 2, 1}
	arrLen := len(arr)

	fmt.Printf("INPUT ARRAY : %+v\n", arr)
	RotateArrayToLeft(arr, arrLen, k)

	arr1 := []int{5, 4, 3, 2, 1}
	arrLen1 := len(arr1)
	RotateArrayToRight(arr1, arrLen1, k)
}

func RotateArrayToLeft(arr []int, arrLen int, k int) {
	//reverse first k elements
	Reverse(arr, 0, k-1)
	//revers remaining last k - n elements
	Reverse(arr, k, arrLen-1)
	//reverse the entire array
	Reverse(arr, 0, arrLen-1)

	fmt.Printf("Rotated array to the left by %v , Arr: %v\n", k, arr)
}

func RotateArrayToRight(arr []int, arrLen int, k int) {
	//reverse first n - k elements
	Reverse(arr, 0, arrLen-k-1)
	//reverse remaining elements n-k elements
	Reverse(arr, arrLen-k, arrLen-1)
	//reverse entire array
	Reverse(arr, 0, arrLen-1)

	fmt.Printf("Rotated array to the right by %v , Arr: %v\n", k, arr)
}
