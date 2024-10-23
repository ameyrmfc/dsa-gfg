package main

import "fmt"

func main() {
	fmt.Println(RotateBy1([]int{1, 2, 3, 4, 5, 6}))
}

func RotateBy1(arr []int) []int {
	arrLen := len(arr)
	temp := arr[0]
	for i := 1; i < arrLen; i++ {
		arr[i-1] = arr[i]
	}
	arr[arrLen-1] = temp
	return arr
}
