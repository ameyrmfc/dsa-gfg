package main

import "fmt"

func main() {
	Permutation([]string{"A", "B", "C"}, 0)
}

//	ABC
//
// ABC BAC CBA
// ABC ACB BAC BCA CBA CAB
func Permutation(arr []string, index int) {
	if index == len(arr)-1 {
		fmt.Print(arr)
		return
	}
	for i := index; i < len(arr); i++ {
		swap(arr, index, i)
		Permutation(arr, index+1)
		swap(arr, index, i)
	}
}

func swap(arr []string, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
