package main

import "fmt"

func main() {
	IsSorted([]int{1, 5, 6, 7, 8, 9})
	IsSorted([]int{11, 1, 5, 6, 7, 8, 9})
}

func IsSorted(arr []int) {
	sorted := true
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			sorted = false
			break
		}
	}
	if sorted {
		fmt.Printf("Array %+v is sorted.\n", arr)
	} else {
		fmt.Printf("Array %+v is not sorted.\n", arr)
	}
}
