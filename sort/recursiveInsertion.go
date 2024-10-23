package sort

import "fmt"

func main() {
	arr := []int{5, 6, 7, 1, 9}
	fmt.Println(RecursiveInsertionSort(arr, len(arr), 0))
}

func RecursiveInsertionSort(arr []int, arrLen int, index int) []int {
	if arrLen-1 == index {
		return arr
	}
	i := index
	for i > 0 && arr[i] < arr[i-1] {
		arr[i], arr[i-1] = arr[i-1], arr[i]
		i--
	}
	return RecursiveInsertionSort(arr, arrLen, index+1)
}
