package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := []int{2, 3, 3, 4, 4, 5, 8}
	fmt.Printf("Input Array1: %+v\nInput Array2: %+v\n------------------------\n", arr1, arr2)
	Union(arr1, arr2)
}

//arr1[] = {2,2,3,4,5}
//arr2[] = {1,2,4,4,6}

func Union(arr, arr2 []int) {
	//opArr
	var output, distinctArr1, distinctArr2, union []int
	//iter on arr1
	i, j := 0, 0

	for i < len(arr) && j < len(arr2) {
		if i > 0 && arr[i] == arr[i-1] {
			i++
			continue
		}
		if j > 0 && arr2[j] == arr2[j-1] {
			j++
			continue
		}
		//if arr[i] < arr[j]; not a common do continue;i++ break to outside loop
		if arr[i] < arr2[j] {
			if len(distinctArr1) == 0 || arr[i] != distinctArr1[len(distinctArr1)-1] && arr[i-1] != arr[i] {
				fmt.Println(distinctArr1, distinctArr2)
				distinctArr1 = append(distinctArr1, arr[i])
			}
			if len(union) == 0 || arr[i] != union[len(union)-1] {
				union = append(union, arr[i])
			}
			i++
			//if arr[i] > arr[j]; not a common do continue;j++;break from current loop
		} else if arr[i] > arr2[j] {
			if len(distinctArr2) == 0 || arr2[j] != distinctArr2[len(distinctArr2)-1] {
				distinctArr2 = append(distinctArr2, arr2[j])
			}
			if len(union) == 0 || arr2[j] != union[len(union)-1] {
				union = append(union, arr2[j])
			}
			j++
			//if arr[i] == arr[j]; created last pushed element and push the value to opArr , then break i++,j++
		} else if arr[i] == arr2[j] {
			if arr2[j] != union[len(union)-1] {
				union = append(union, arr[i])
			}
			if len(output) == 0 || output[len(output)-1] != arr[i] {
				output = append(output, arr[i])
			}
			i++
			j++
		}
	}
	for k := i; k < len(arr); k++ {
		union = append(union, arr[k])
		distinctArr1 = append(distinctArr1, arr[k])
	}
	for k := j; k < len(arr2); k++ {
		union = append(union, arr2[k])
		distinctArr2 = append(distinctArr2, arr2[k])
	}
	fmt.Printf("common elemetns : %+v\n", output)
	fmt.Printf("distinct elemetns arrray 1: %+v\n", distinctArr1)
	fmt.Printf("distinct elemetns array 2: %+v\n", distinctArr2)
	fmt.Printf("union is: %+v\n", union)
}
