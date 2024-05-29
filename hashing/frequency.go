package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 1, 1, 1, 3, 4}
	FrequencyOfNo(arr)
}

func FrequencyOfNo(arr []int) {
	freq := make(map[int]int)
	for _, val := range arr {
		if _, ok := freq[val]; ok {
			freq[val]++
		} else {
			freq[val] = 1
		}
	}
	fmt.Printf("Frequency : %+v\n", freq)
}
