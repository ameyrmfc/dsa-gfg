package main

import "fmt"

func main() {
	arr := []int{6, 5, 4, 3, 2, 1}
	fmt.Println(Insertion(arr))
}

func Insertion(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			temp := arr[j-1]
			arr[j-1] = arr[j]
			arr[j] = temp
			j--
		}
	}
	return arr
}
