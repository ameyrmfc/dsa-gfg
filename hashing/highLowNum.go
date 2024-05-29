package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 1, 1, 1, 3, 4}
	HighLowNo(arr)
}

func HighLowNo(arr []int) {
	freq := make(map[int]int)
	var higNum, highFreq, lowNum, lowFreq int = 0, 0, 0, len(arr)
	for _, val := range arr {
		if _, ok := freq[val]; ok {
			freq[val]++
		} else {
			freq[val] = 1
		}
	}
	for key, value := range freq {
		if value >= highFreq {
			higNum = key
			highFreq = value
		}

		if value <= lowFreq {
			lowNum = key
			lowFreq = value
		}
	}
	fmt.Printf("High Freq No: %v , Frequecy: %v\n", higNum, highFreq)
	fmt.Printf("Low Freq No: %v , Frequecy: %v\n", lowNum, lowFreq)
}
