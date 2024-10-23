package sort

//move the smallest element to the left
//keep on comparing smallest and moving it to the left most index
import "fmt"

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println(SelectionSort(arr))
}

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		small := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				small = j
			}
		}
		if i != small {
			arr[i], arr[small] = arr[small], arr[i]
		}
		fmt.Printf("%v\n", arr)
	}
	return arr
}
