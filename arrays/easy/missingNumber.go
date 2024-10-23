package main

import "fmt"

func main() {
	missingNumber([]int{1, 3}, 3)
}

func missingNumber(arr []int, n int) {
	sum1 := (n * (n + 1)) / 2
	sum2 := 0
	for _, value := range arr {
		sum2 += value
	}
	fmt.Printf("Missing Number in Array %+v  is %v\n", arr, sum1-sum2)
}
