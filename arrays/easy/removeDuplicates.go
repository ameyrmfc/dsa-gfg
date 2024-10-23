package main

import (
	"fmt"
	"time"
)

func main() {
	ip := make([]int, 0)
	for i := 0; i < 10000000; i++ {
		ip = append(ip, i)
	}

	st := time.Now()
	RemoveDuplicatesMapIterate(ip)
	fmt.Println(time.Since(st))

	st1 := time.Now()
	RemoveDuplicates(ip)
	fmt.Println(time.Since(st1))
}

func RemoveDuplicates(arr []int) []int {
	cacheMap := make(map[int]int)
	temp := 0
	for i := 0; i < len(arr); i++ {
		value := arr[i]
		if _, exist := cacheMap[value]; exist {
			continue
		} else {
			cacheMap[value] = 1
			arr[temp] = value
			temp++
		}
	}
	return arr
}
func RemoveDuplicatesMapIterate(arr []int) []int {
	op := make([]int, 0)
	cacheMap := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		value := arr[i]
		if _, exist := cacheMap[value]; exist {
			continue
		} else {
			cacheMap[value] = 1
		}
	}
	for key, _ := range cacheMap {
		op = append(op, key)
	}
	return op
}
