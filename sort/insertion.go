package sort

//NOTE use the same array and keep on moving and compare with all previous elements
import "fmt"

func main() {
	arr := []int{6, 4, 3, 2, 1, 5}
	fmt.Println(Insertion(arr))
}

// {6, 4, 3, 2, 1, 5}
// {6, 4, 3, 2, 1, 5}
func Insertion(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			temp := arr[j-1]
			arr[j-1] = arr[j]
			arr[j] = temp
			j--
		}
	}
	return arr
}
