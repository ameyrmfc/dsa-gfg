package main

import "fmt"

func main() {
	input := []int{2, 0, 2, 1, 1, 0}
	//fmt.Println(Sort012SMerge(input))
	fmt.Println(Sort012S(input))
}

func Sort012S(arr []int) []int {
	low := 0
	mid := 0
	high := len(arr) - 1
	for i := 0; i < len(arr); i++ {
		if arr[mid] == 0 {
			arr[low], arr[mid] = arr[mid], arr[low]
			mid++
			low++
		} else if arr[mid] == 1 {
			mid++
		} else if arr[mid] == 2 {
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}
	return arr
}

func Sort012SMerge(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := Sort012SMerge(arr[:mid])
	right := Sort012SMerge(arr[mid:])
	return MergeSort(left, right)
}

func MergeSort(left, right []int) []int {
	i := 0
	j := 0
	var op []int
	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			op = append(op, right[j])
			j++
		} else {
			op = append(op, left[i])
			i++
		}
	}
	for k := i; k < len(left); k++ {
		op = append(op, left[k])
	}
	for k := j; k < len(right); k++ {
		op = append(op, right[k])
	}
	return op
}
