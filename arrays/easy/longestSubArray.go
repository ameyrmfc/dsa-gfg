package main

import "fmt"

func main() {
	//longestSubArray([]int{6, 2, 3, 1, 1, 1, 1, 1, 1}, 6)
	//longestSubArrayLength([]int{6, 2, 3, 1, 1, 1, 1, 1, 1}, 6)
	//longestSubArray([]int{-1, 1, 1, 1, 1, 1}, 3)
	//longestSubArrayLength([]int{-1, 1, 1, 1, 1, 1}, 3)
	longestSubArray([]int{1, 2, 1, 2, 1}, 3)
}

func longestSubArray(arr []int, k int) {
	startIndex := 0
	var output []int
	sum := arr[0]
	i := 1
	for i < len(arr) {
		sum = arr[i] + sum
		if sum >= k {
			//if sum is matching ,and length of current existing op is less
			//than current one then replace the array
			//i-startIndex [0,1,2,3,4,5] if i=5 and startIndex=2 , then the array length will be 3
			if sum == k && len(output) < (i-startIndex)+1 {
				output = arr[startIndex : i+1]
			}
			startIndex += 1
			i = startIndex
			sum = 0
		} else {
			i++
		}
	}
	fmt.Printf("\nInput : %+v , sum : %+v \t Longest Subarray: %+v\n", arr, k, output)
}

func longestSubArrayLength(arr []int, k int) {
	startIndex := 0
	var output int
	sum := arr[0]
	i := 1
	for i < len(arr) {
		sum = arr[i] + sum
		if sum >= k {
			//if sum is matching ,and length of current existing op is less
			//than current one then replace the array
			//i-startIndex [0,1,2,3,4,5] if i=5 and startIndex=2 , then the array length will be 3
			if sum == k && output < (i-startIndex)+1 {
				output = (i - startIndex) + 1
			}
			startIndex += 1
			i = startIndex
			sum = 0
		} else {
			i++
		}
	}
	fmt.Printf("\nInput : %+v , sum : %+v \t Longest Subarray: %+v\n", arr, k, output)
}
