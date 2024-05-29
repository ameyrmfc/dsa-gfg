package main

import "fmt"

//Input: N = 5
//Output: 0 1 1 2 3 5

func main() {
	n := 5
	//fibonacci(n)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", RecursiveFib(i))
	}
}

func fibonacci(n int) {
	num1, num2 := 0, 1
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", num1)
		num3 := num1 + num2
		num1 = num2
		num2 = num3
		fmt.Println(num3)
	}
}

func RecursiveFib(n int) int {
	if n <= 1 {
		return n
	}
	return RecursiveFib(n-1) + RecursiveFib(n-2)
}
