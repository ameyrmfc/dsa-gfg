package main

import (
	"fmt"
	"slices"
)

func main() {
	input := []int{1, 2, 3, 4, 5, 6}
	k := 9
	//arr, str := Find2sumMap(input, k)
	//fmt.Printf("Input : %+v , %v %v\n", input, arr, str)
	fmt.Println(Find2sumSort(input, k))
}

func Find2sumMap(arr []int, sum int) ([2]int, string) {
	cacheMap := make(map[int]int)
	for index, value := range arr {
		searchEle := sum - value
		if arrValue, ok := cacheMap[searchEle]; ok {
			return [2]int{arrValue, index}, "YES"
		} else {
			cacheMap[value] = index
		}
	}
	return [2]int{-1, -1}, "NO"
}

func Find2sumSort(arr []int, sum int) ([]int, string) {
	slices.Sort(arr)
	left := 0
	right := len(arr) - 1
	for left < right {
		arraySum := arr[left] + arr[right]
		if arraySum < sum {
			left++
		} else if arraySum > sum {
			right--
		} else if arraySum == sum {
			return []int{left, right}, "YES"
		}
	}
	return []int{-1, -1}, "NO"
}
