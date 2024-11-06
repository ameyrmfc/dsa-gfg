package main

import "fmt"

func main() {
	input := []int{4, 7, 1, 0}
	op := LeadersOfArrayOpt(input)
	fmt.Printf("Input: %+v Leaders Of Arrays: %+v\n", input, op)
}

func LeadersOfArray(arr []int) []int {
	var op []int
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			op = append(op, arr[i-1])
		} else if arr[i-1] < arr[i] {
			for j := len(op) - 1; j >= 0; j-- {
				if op[j] > arr[i] {
					break
				}
				op = op[:len(op)-1]
			}
		}
	}
	op = append(op, arr[len(arr)-1])
	return op
}

// [10, 22, 12, 3, 0, 6]
func LeadersOfArrayOpt(arr []int) []int {
	var op []int
	maxLen := len(arr) - 1
	maxEle := arr[maxLen]
	op = append(op, maxEle)

	for i := maxLen - 1; i >= 0; i-- {
		if arr[i] > maxEle {
			op = append([]int{arr[i]}, op...)
			maxEle = arr[i]
		}
	}
	return op
}
