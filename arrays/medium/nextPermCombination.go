package main

import "fmt"

func main() {
	input := []string{"A", "B", "C"}
	NextPermutation(input, input, 0, false)
}

//	ABC
//
// ABC BAC CBA
// ABC ACB BAC BCA CBA CAB
func NextPermutation(arr, input []string, index int, flag bool) {
	if index == len(arr)-1 {
		if !flag && arr == input {

		}
		fmt.Print(arr)
		return
	}
	for i := index; i < len(arr); i++ {
		NextSwap(arr, index, i)
		NextPermutation(arr, input, index+1, flag)
		NextSwap(arr, index, i)
	}
}

func NextSwap(arr []string, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
