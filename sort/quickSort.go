package sort

import "fmt"

func main() {
	arr := []int{9, 6, 5, 4, 1}
	//RecursiveQuickSort(arr, 0, len(arr)-1)
	RQSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func RecursiveQuickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		RecursiveQuickSort(arr, low, pi-1)
		RecursiveQuickSort(arr, pi+1, high)
	}
}

func QuickSort(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := i + 1; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[i], arr[j]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func RQSort(arr []int, low, high int) {
	if low < high {
		pi := QuickSort(arr, low, high)
		RQSort(arr, low, pi-1)
		RQSort(arr, pi+1, high)

	}
}
