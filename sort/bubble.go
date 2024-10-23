package sort

//NOTE HERE we have to move the highest element to right most index and repeat the process
import "fmt"

func main() {
	arr := []int{5, 4, 3, 2, 6, 1}
	fmt.Println(BubbleSort(arr))
}

func BubbleSort(arr []int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}
