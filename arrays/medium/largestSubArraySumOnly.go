package main

import "fmt"

func main() {
	LargestSumSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	LargestSumSubArrayOptimized([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
}

func LargestSumSubArray(arr []int) int {
	maxSum := 0
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if sum > maxSum {
				maxSum = sum
			}
		}

	}
	fmt.Printf("The maximum subarray sum is: %v\n", maxSum)
	return maxSum
}

func LargestSumSubArrayOptimized(arr []int) int {
	maxSum := 0
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if sum > maxSum {
			maxSum = sum
		}

		if sum < 0 {
			sum = 0
		}
	}
	fmt.Printf("The maximum subarray Optimized sum is: %v\n", maxSum)
	return maxSum
}
