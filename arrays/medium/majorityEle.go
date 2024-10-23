package main

import "fmt"

func main() {
	fmt.Println(FindMajority([]int{2, 1, 2, 2, 2, 2, 3, 3}))
}

func FindMajority(arr []int) int {
	n := len(arr)
	cnt := 0
	el := 0

	for i := 0; i < n; i++ {
		if cnt == 0 {
			cnt = 1
			el = arr[i]
		} else if el == arr[i] {
			cnt++
		} else {
			cnt--
		}
	}

	cnt1 := 0
	for i := 0; i < n; i++ {
		if arr[i] == el {
			cnt1++
		}
	}

	if cnt1 > (n / 2) {
		return el
	}
	return -1
}
