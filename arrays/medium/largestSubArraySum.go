package main

import "fmt"

func main() {
	LargestSumSubArraySlice([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
}

func LargestSumSubArraySlice(arr []int) int {
	maxSum := 0
	sum := 0
	startIndex := 0
	endIndex := 0
	start := 0
	for i := 0; i < len(arr); i++ {
		if sum == 0 {
			start = i
		}
		sum += arr[i]

		if sum > maxSum {
			startIndex = start
			endIndex = i
			maxSum = sum
		}

		if sum < 0 {
			sum = 0
		}
	}
	fmt.Printf("The maximum subarray Optimized sum is: %v, Slice: %+v\n ", maxSum, arr[startIndex:endIndex+1])
	return maxSum
}
