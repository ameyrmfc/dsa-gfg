package main

import "fmt"

func main() {
	arr := []int{0, 1, 0, 2, 0, 3, 0, 5}
	fmt.Printf("INPUT: %+v\n", arr)
	MoveZerosToEnd(arr)
	fmt.Printf("Output: %+v\n", arr)

}

// [1,0,2,3,0,5]
// [1,2,0,3,0,5] zindex:= 2,i = 2
// [1,2,3,0,0,5] zindex:= 2,i = 3
func MoveZerosToEnd(arr []int) {
	zeroIndex := 0
	for i := 1; i < len(arr); i++ {
		zeroVal := arr[zeroIndex]
		indexVal := arr[i]
		if zeroVal == 0 && indexVal > 0 {
			arr[i], arr[zeroIndex] = arr[zeroIndex], arr[i]
			zeroIndex += 1
		} else if zeroVal != 0 && indexVal == 0 {
			zeroIndex = i
		}
	}
}
