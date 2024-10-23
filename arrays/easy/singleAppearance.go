package main

import "fmt"

func main() {
	findSingleAppearanceXor([]int{1, 2, 2, 3, 3, 4, 4, 5, 5})
}

func findSingleAppearanceXor(arr []int) {
	xor := 0
	for _, value := range arr {
		xor = xor ^ value
	}
	fmt.Printf("Input: %+v \nNumber which appered once is %v\n", arr, xor)
}
