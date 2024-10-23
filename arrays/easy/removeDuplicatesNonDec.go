package main

import (
	"fmt"
	"time"
)

func main() {
	ip := []int{1, 1, 2, 3, 4, 4, 5}
	st := time.Now()
	fmt.Println(RemoveDuplicatesDec(ip))
	fmt.Println(time.Since(st))
}

func RemoveDuplicatesDec(arr []int) []int {
	temp := 0
	for i := temp + 1; i < len(arr); i++ {
		if arr[temp] != arr[i] {
			temp += 1
			arr[temp] = arr[i]
		}
	}
	return arr
}
