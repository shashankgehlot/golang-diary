package main

import (
	"fmt"
)

func selection_sort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		fmt.Println("min", i)
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				fmt.Println("min chage", j)
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}

func main() {
	array := []int{10, 9, 8, 7, 30, 20, 21, 31, 22, 55, 66}
	fmt.Println(array)
	fmt.Println(selection_sort(array))

}
