package main

import "fmt"

func main() {
	RearrangePosNegArray([]int{1, 2, -4, -5})
}

func RearrangePosNegArray(arr []int) {
	pos := 0
	neg := 1
	op := make([]int, len(arr))
	for _, value := range arr {
		if value < 0 {
			op[neg] = value
			neg += 2
		} else {
			op[pos] = value
			pos += 2
		}
	}
	fmt.Printf("Input :%+v , Rearrange Output: %+v\n", arr, op)
}
