package main

import "fmt"

func main() {
	arr := []int{1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1}
	fmt.Printf("Input Array :%+v, Consecutive ones: %v\n", arr, consecutiveOnes(arr))
}

func consecutiveOnes(arr []int) int {
	maxConOnes := 0
	temp := 0
	for _, value := range arr {
		if value != 1 {
			temp = 0
		} else {
			temp++
		}
		if maxConOnes < temp {
			maxConOnes = temp
		}
	}
	return maxConOnes
}
