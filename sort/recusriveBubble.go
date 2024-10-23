package sort

import "fmt"

func main() {
	arr := []int{5, 6, 7, 1, 9}
	fmt.Println(RecursiveBubbleSort(arr, len(arr)))
}

func RecursiveBubbleSort(arr []int, arrLen int) []int {
	if arrLen == 0 {
		return arr
	}
	for i := 0; i < arrLen-1; i++ {
		if arr[i] > arr[arrLen-1] {
			arr[i], arr[arrLen-1] = arr[arrLen-1], arr[i]
		}
	}
	return RecursiveBubbleSort(arr, arrLen-1)
}
