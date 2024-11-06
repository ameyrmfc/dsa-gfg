package main

import (
	"fmt"
	"math"
)

func main() {
	BuySellStocksCond([]int{7, 6, 4, 3, 1})
	BuySellStocksCond([]int{7, 1, 5, 3, 6, 4})
}

func BuySellStocks(arr []int) {
	buy := float64(arr[0])
	profit := float64(0)
	for i := 0; i < len(arr); i++ {
		buy = math.Min(buy, float64(arr[i]))
		profit = math.Max(profit, float64(arr[i])-buy)
	}
	fmt.Printf("Profit from stock is : %+v\n", profit)
}

func BuySellStocksCond(arr []int) {
	buy := arr[0]
	profit := 0
	for i := 0; i < len(arr); i++ {
		if buy > arr[i] {
			buy = arr[i]
		}
		if profit < arr[i]-buy {
			profit = arr[i] - buy
		}
	}
	fmt.Printf("Profit from stock is : %+v\n", profit)
}
