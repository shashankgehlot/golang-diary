package main

import (
	"fmt"
)

func sorted(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
	}
	return a
}

func binary_search(arr []int, val int) bool {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		fmt.Println("low>>>", low, "high>>>>", high, "mid>>>>", mid)
		if arr[mid] == val {
			return true
		}
		if arr[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if low == len(arr) || arr[low] != val {
		return false
	}
	return true
}

func main() {
	array := []int{10, 9, 8, 7, 30, 20, 21, 31, 22, 55, 66}
	fmt.Println("array", array)
	sorted_array := sorted(array)
	fmt.Println("sorted_array", sorted_array)
	value := 55
	sg := binary_search(sorted_array, value)
	fmt.Println(sg)
}

// func binarySearch(needle int, haystack []int) bool {

// 	low := 0
// 	high := len(haystack) - 1

// 	for low <= high {
// 		fmt.Println("low", low, "high", high)
// 		median := (low + high) / 2

// 		if haystack[median] < needle {
// 			low = median + 1
// 		} else {
// 			high = median - 1
// 		}
// 	}

// 	if low == len(haystack) || haystack[low] != needle {
// 		return false
// 	}

// 	return true
// }

// func main() {
// 	items := []intc
// 	fmt.Println(binarySearch(28, items))
// }
